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

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
