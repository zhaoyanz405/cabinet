package main

import (
	"cabinet/benchmark/cacheClient"
	"flag"
	"fmt"
)

func main() {
	server := flag.String("h", "localhost", "cache server address")
	op := flag.String("c", "get", "command, could be get/set/del")
	key := flag.String("k", "", "key")
	value := flag.String("v", "", "value")
	flag.Parse()
	client := cacheClient.New("tcp", *server)
	cmd := &cacheClient.Cmd{Name: *op, Key: *key, Value: *value}
	client.Run(cmd)
	if cmd.Error != nil {
		fmt.Println("error:", cmd.Error)
	} else {
		fmt.Println(cmd.Value)
	}
}
