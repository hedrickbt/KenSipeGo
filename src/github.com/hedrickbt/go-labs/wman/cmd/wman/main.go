package main

import (
	//"fmt"
	//"github.com/hedrickbt/go-labs/wman/pkg/string"
	"github.com/hedrickbt/go-labs/wman/pkg/cmd"
	"os"
)

func main() {
	//arg := os.Args[1]
	//fmt.Println("hello", arg)
	//fmt.Printf("hello %s, your name backward is %q", arg, string.Reverse(arg))
	if err := cmd.NewWmanCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
