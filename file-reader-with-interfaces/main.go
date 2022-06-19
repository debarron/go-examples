package main

import (
	"os"
	"io"
)

func main(){
	args := os.Args[1:]
	fileName := args[0]

	reader, _ := os.Open(fileName)
	io.Copy(os.Stdout, reader)
}
