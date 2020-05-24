package main

import (
	"fmt"
	"io"
	"net/http"

	// "io"

	// "net/http"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type challenge struct {
	link string
	name string
}

var tabs int

func main() {
	tabs = 0
	getChllanges()
}

func getPage() io.ReadCloser {
	resp, _ := http.Get("https://www.hackerearth.com/challenges/")
	// dat, _ := ioutil.ReadFile("template.html")
	// file,_ :=
	return resp.Body
	// return bytes.NewReader(dat)
}

// func getPage() *bytes.Reader {
// 	// resp, _ := http.Get("https://www.hackerearth.com/challenges/")
// 	dat, _ := ioutil.ReadFile("template.html")
// 	return bytes.NewReader(dat)
// }

func bfs(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	t := n.FirstChild
	for t != nil {
		if t.DataAtom == atom.Div {
			for j := 0; j < len(t.Attr); j++ {
				if t.Attr[j].Key == "class" && t.Attr[j].Val == "ongoing challenge-list" {
					return t
				}
			}
		}
		if bfs(t) != nil {
			return bfs(t)
		}
		t = t.NextSibling
	}
	return nil
}

func getChallenge(n *html.Node) challenge {
	var chal challenge
	chal.link = n.Attr[2].Val
	chal.name = n.FirstChild.NextSibling.FirstChild.NextSibling.Attr[1].Val
	return chal
}

func getChllanges() []challenge {
	htm, _ := html.Parse(getPage())
	challenges := make([]challenge, 0)
	challengesNode := bfs(htm)
	t := challengesNode.FirstChild
	for t != nil {
		if len(t.Attr) != 0 && t.Attr[0].Key == "class" && t.Attr[0].Val == "challenge-card-modern" {
			// fmt.Println(t.FirstChild.NextSibling.Attr)
			challenges = append(challenges, getChallenge(t.FirstChild.NextSibling))
		}
		t = t.NextSibling
	}
	for _, ch := range challenges {
		fmt.Println(ch.name + " : " + ch.link + " | ")
	}

	return challenges
}
