package main

import (
	"github.com/gdey/loglog/foo"
	"github.com/gdey/loglog/log"
)

func init() {
	log.SetBasePath()
}

func main() {
	log.WithContext("Gautam").Printf(" err %v: %v", 56, "this is error 56")
	log.WithContext("Gautam").Println(" err %v: %v", 56, "this is error 56")
	foo.Foo()
}
