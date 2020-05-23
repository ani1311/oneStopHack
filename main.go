package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

type challenge struct {
	link string
	name string
}

func main() {
	getChllanges()
}

func getPage() io.ReadCloser {
	resp, _ := http.Get("https://www.hackerearth.com/challenges/")
	// file,_ :=
	return resp.Body
}

func bfs(n *html.Node) {
	if n == nil {
		return
	}
	t := n.FirstChild
	for t != nil {
		fmt.Print(t.Data + " ")
		bfs(t)
		t = t.NextSibling
	}
}

func getChllanges() []challenge {
	htm, _ := html.Parse(getPage())
	challenges := make([]challenge, 10)
	bfs(htm)
	return challenges
}
