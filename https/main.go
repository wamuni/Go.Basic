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
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// fmt.Println(resp.Status)
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Proto)
	// response_body := make([]byte, 9999)
	// n, err := resp.Body.Read(response_body)
	// if err != nil {
	// 	fmt.Print("Error: ", err)
	// }
	// fmt.Println("number returns: ", n)
	// fmt.Println(string(response_body))

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

// now the logWriter implements Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
