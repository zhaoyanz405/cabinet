package main

import (
	"cabinet/cache"
	"cabinet/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
