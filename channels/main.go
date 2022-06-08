package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://github.com/erikseyti",
		"http://google.com.br",
		"https://www.linkedin.com/in/erik-seyti/",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	// alternative way to loop a infinite a channel
	for l := range c {
		
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
