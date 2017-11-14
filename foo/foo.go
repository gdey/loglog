package foo

import (
	"github.com/gdey/loglog/log"
)

func Foo() {
	log.WithContext("Gautam").Printf(" err %v: %v", 56, "this is error 56")
	log.WithContext("Gautam").Println(" err %v: %v", 56, "this is error 56")
}
