package main

import (
	"cabinet/cache"
	"cabinet/server"
)


func main() {
	c := cache.New("inmemory")
	server.New(c).Listen()
}