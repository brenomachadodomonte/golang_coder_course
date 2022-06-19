package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// readonly channels
func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			if len(r.FindStringSubmatch(string(html))) > 1 {
				c <- r.FindStringSubmatch(string(html))[1]
			} else {
				c <- "Title not found"
			}

		}(url)
	}

	return c
}

func forward(origin <-chan string, recipient chan string) {
	for {
		recipient <- <-origin
	}
}

func join(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(input1, c)
	go forward(input2, c)

	return c
}

func main() {
	c := join(
		title("https://brenomachado.dev", "https://www.google.com"),
		title("https://www.youtube.com", "https://amazon.com.br"),
	)

	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}
