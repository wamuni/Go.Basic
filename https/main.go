package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
	response_body := make([]byte, 9999)
	n, err := resp.Body.Read(response_body)
	if err != nil {
		fmt.Print("Error: ", err)
	}
	fmt.Println("number returns: ", n)
	fmt.Println(string(response_body))
}
