package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		return
	}

	fmt.Println(url, "is up!")
}
