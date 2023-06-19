package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error occured : {}", err)
		os.Exit(1)
	}

	// fmt.Println("response from browser :{}", resp)

	//this make is an inbuilt function which takes a type and default value to which it will initialize the type with empty spaces
	//make function only takes slice,map or chan as type and it returns same object type and not a pointer to that object
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// this line of code basically do the same thing which we are doing above
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Printed this much of bytes : ", len(bs))
	return len(bs), nil
}
