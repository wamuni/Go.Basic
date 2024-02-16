package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// how to create channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// infinite loop
	for l := range c {
		go func(url string) {
			time.Sleep(3 * time.Second)
			checkLink(url, c)
		}(l)
	}
}

// when use channel in parameter, we have to specify data type as well
func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
