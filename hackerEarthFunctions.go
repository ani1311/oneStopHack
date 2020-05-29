package main

import (
	"fmt"
	"unicode"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getChllanges() []Challenge {
	htm, _ := html.Parse(getPage("https://www.hackerearth.com/challenges/"))
	challenges := make([]Challenge, 0)
	challengesNode := getChallengeNode(htm, atom.Div, "class", "ongoing challenge-list")
	t := challengesNode.FirstChild
	for t != nil {
		if len(t.Attr) != 0 && t.Attr[0].Key == "class" && t.Attr[0].Val == "challenge-card-modern" {
			chal, err := getChallenge(t.FirstChild.NextSibling)
			if err != nil {
				fmt.Println("challenge skipped")
			} else {
				challenges = append(challenges, chal)

			}
		}
		t = t.NextSibling
	}
	// for _, ch := range challenges {
	// 	fmt.Println(ch.name + " : " + ch.link + " | ")
	// }

	return challenges
}

func getChallenge(n *html.Node) (Challenge, error) {
	var chal Challenge
	chal.Link = n.Attr[2].Val
	chal.Name = n.FirstChild.NextSibling.FirstChild.NextSibling.Attr[1].Val
	if unicode.IsLower(rune(chal.Name[0])) {
		return chal, fmt.Errorf("BAD")
	}
	return chal, nil
}
