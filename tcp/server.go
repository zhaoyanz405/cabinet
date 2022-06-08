package tcp

import (
	"bufio"
	"cabinet/cache"
	"cabinet/cluster"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	cache.Cache
	cluster.Node
}

func (s *Server) Listen() {
	l, e := net.Listen("tcp", ":12346")
	if e != nil {
		panic(e)
	}
	log.Println("tcp ready to serve")

	for {
		c, e := l.Accept()
		if e != nil {
			panic(e)
		}

		go s.process(c)
	}
}

func New(c cache.Cache, n cluster.Node) *Server {
	return &Server{c, n}
}

func (s *Server) readKey(r *bufio.Reader) (string, error) {
	klen, e := readLen(r)
	if e != nil {
		return "", e
	}

	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", e
	}

	key := string(k)
	addr, ok := s.ShouldProcess(key)
	if !ok {
		return "", errors.New("redirect " + addr)
	}

	return key, nil
}

func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	klen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}

	vlen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}

	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", nil, e
	}

	key := string(k)
	addr, ok := s.ShouldProcess(key)
	if !ok {
		return "", nil, errors.New("redirect " + addr)
	}

	v := make([]byte, vlen)
	_, e = io.ReadFull(r, v)
	if e != nil {
		return "", nil, e
	}

	return key, v, nil
}

func readLen(r *bufio.Reader) (int, error) {
	tmp, e := r.ReadString(' ')
	if e != nil {
		return 0, e
	}

	l, e := strconv.Atoi(strings.TrimSpace(tmp))
	if e != nil {
		return 0, e
	}

	return l, nil
}

func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("-%d ", len(errString)) + errString
		_, e := conn.Write([]byte(tmp))
		return e
	}

	vlen := fmt.Sprintf("%d ", len(value))
	_, e := conn.Write(append([]byte(vlen), value...))
	return e
}

func (s *Server) get(ch chan chan *result, r *bufio.Reader) {
	c := make(chan *result)
	ch <- c
	k, e := s.readKey(r)
	if e != nil {
		c <- &result{nil, e}
		return
	}

	go func() {
		v, e := s.Get(k)
		c <- &result{v, e}
	}()
}

func (s *Server) set(ch chan chan *result, r *bufio.Reader) {
	c := make(chan *result)
	ch <- c

	k, v, e := s.readKeyAndValue(r)
	if e != nil {
		c <- &result{nil, e}
		return
	}

	go func() {
		c <- &result{nil, s.Set(k, v)}
	}()
}

func (s *Server) del(ch chan chan *result, r *bufio.Reader) {
	c := make(chan *result)
	ch <- c

	k, e := s.readKey(r)
	if e != nil {
		c <- &result{nil, e}
		return
	}
	go func() {
		c <- &result{nil, s.Del(k)}
	}()
}

type result struct {
	v []byte
	e error
}

func (s *Server) process(conn net.Conn) {
	r := bufio.NewReader(conn)
	resultCh := make(chan chan *result, 5000)
	defer close(resultCh)

	go reply(conn, resultCh)
	for {
		op, e := r.ReadByte()
		if e != nil {
			if e != io.EOF {
				log.Println("close connection due to error:", e)
			}
			return
		}
		if op == 'S' {
			s.set(resultCh, r)
		} else if op == 'G' {
			s.get(resultCh, r)
		} else if op == 'D' {
			s.del(resultCh, r)
		} else {
			log.Println("close connection due to invalid operation:", op)
			return
		}
	}
}

func reply(conn net.Conn, resultCh chan chan *result) {
	defer conn.Close()
	for {
		c, open := <-resultCh
		if !open {
			return
		}
		r := <-c
		e := sendResponse(r.v, r.e, conn)
		if e != nil {
			log.Println("close connection due to error:", e)
			return
		}
	}
}
