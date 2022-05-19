package main

import (
	"fmt"
	"github.com/hedrickbt/go-labs/wman/pkg/string"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Println("hello", arg)
	fmt.Printf("hello %s, your name backward is %q", arg, string.Reverse(arg))
}
