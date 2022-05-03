package main

import (
	"cabinet/cache"
	server "cabinet/tcp"
)

func main() {
	c := cache.New("inmemory")
	server.New(c).Listen()
}
