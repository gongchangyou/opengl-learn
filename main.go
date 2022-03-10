package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func load() {

}

func main() {
	fmt.Println("hello world")
}
