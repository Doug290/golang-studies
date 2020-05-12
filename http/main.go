package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//// what is done here is that we create a large enought byte slice to pass to the
	//// Read() interface, that will fill that byte slice with the http response body
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
