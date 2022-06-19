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
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := title("https://brenomachado.dev", "https://www.google.com")
	t2 := title("https://www.youtube.com")

	fmt.Println("Firsts Ones: ", <-t1, " | ", <-t2)
	fmt.Println("Seconds Ones: ", <-t1, " | ")
}
