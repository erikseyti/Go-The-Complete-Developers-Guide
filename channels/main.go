package main

import (
	"fmt"
	"net/http"
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

	for i :=0; i< len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up!"
}
