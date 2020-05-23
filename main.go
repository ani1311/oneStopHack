package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	"bytes"
	// "net/http"
	"golang.org/x/net/html"
)

type challenge struct {
	link string
	name string
}
var tabs int =0
func main() {
	getChllanges()
}

func getPage() *bytes.Reader {
	//resp, _ := http.Get("https://www.hackerearth.com/challenges/")
	dat, _ := ioutil.ReadFile("template.html")
	//file,_ :=
	// return resp.Body
	return bytes.NewReader(dat)
}

func bfs(n *html.Node) {
	if n == nil {
		return
	}
	t := n.FirstChild
	tabs = tabs +1
	for t != nil {
		for i := 0; i < tabs; i++ { 
			fmt.Print("\t")
		}

		fmt.Println(t.Data)
		bfs(t)
		t = t.NextSibling

	}
	tabs = tabs -1
}

func getChllanges() []challenge {
	htm, _ := html.Parse(getPage())
	challenges := make([]challenge, 10)
	bfs(htm)
	return challenges
}
