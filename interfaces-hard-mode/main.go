package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, f)
	// bs := make([]byte, 24*1024)
	//f.Read(bs)

}
