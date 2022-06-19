package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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

func fastest(url1, url2, url3 string) string {
	c1 := title(url1)
	c2 := title(url2)
	c3 := title(url3)

	// control for concorrency
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(2 * time.Second):
		return "Everyone loose"
	}
}

func main() {
	champion := fastest(
		"https://www.google.com",
		"https://www.youtube.com",
		"https://amazon.com.br",
	)

	fmt.Println(champion)
}
