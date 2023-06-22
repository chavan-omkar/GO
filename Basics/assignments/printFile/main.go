package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fileName := os.Args[1]

	resp, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error Occurred :{}", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp)
}
