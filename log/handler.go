package log

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

var truncPath string

func SetBasePath() {
	var pc [1]uintptr
	runtime.Callers(2, pc[:])
	fn := runtime.FuncForPC(pc[0])
	ffile, _ := fn.FileLine(pc[0])
	truncPath = filepath.Dir(ffile)
}

func caller() (string, int, string) {
	var pc [1]uintptr
	var err error
	runtime.Callers(3, pc[:])
	fn := runtime.FuncForPC(pc[0])
	ffile, ln := fn.FileLine(pc[0])
	file := ffile
	//	file := filepath.Base(ffile)
	if truncPath != "" {
		file, err = filepath.Rel(truncPath, ffile)
		if err != nil {
			file = filepath.Base(ffile)
		}
	}
	return file, ln, fn.Name()
}

func Printf(prefix string, format string, values ...interface{}) {
	msg := fmt.Sprintf(format, values...)
	file, ln, fn := caller()
	log.Printf("%v:%v(%v) [ctx: %v] %v", file, ln, fn, prefix, msg)
}

func Println(prefix string, values ...interface{}) {
	msg := fmt.Sprintln(values...)
	file, ln, fn := caller()
	log.Printf("%v:%v(%v) [ctx: %v] %v", file, ln, fn, prefix, msg)
}

type Log string

func (l Log) Printf(format string, values ...interface{}) {
	msg := fmt.Sprintf(format, values...)
	file, ln, fn := caller()
	log.Printf("%v:%v(%v) [ctx: %v] %v", file, ln, fn, string(l), msg)
}
func (l Log) Println(values ...interface{}) {
	msg := fmt.Sprintln(values...)
	file, ln, fn := caller()
	log.Printf("%v:%v(%v) [ctx: %v] %v", file, ln, fn, string(l), msg)
}

func WithContext(ctx string) Log {
	return Log(ctx)
}
