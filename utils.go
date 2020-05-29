package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getPage(page string) io.ReadCloser {
	resp, _ := http.Get(page)
	return resp.Body
}

func getPageFromLocal(page string) io.ReadCloser {
	dat, _ := ioutil.ReadFile(page)
	return ioutil.NopCloser(bytes.NewReader(dat))
}

func getChallengeNode(n *html.Node, atomType atom.Atom, attribKey string, attribVal string) *html.Node {
	if n == nil {
		return nil
	}
	t := n.FirstChild
	for t != nil {
		if t.DataAtom == atomType {
			for j := 0; j < len(t.Attr); j++ {
				if t.Attr[j].Key == attribKey && t.Attr[j].Val == attribVal {
					return t
				}
			}
		}
		chalNode := getChallengeNode(t, atomType, attribKey, attribVal)
		if chalNode != nil {
			return chalNode
		}
		t = t.NextSibling
	}
	return nil
}

// func getChallengeNode(n *html.Node, atomType atom.Atom, tagVal string) *html.Node {
// 	if n == nil {
// 		return nil
// 	}
// 	t := n.FirstChild
// 	for t != nil {
// 		if t.DataAtom == atomType {
// 			for j := 0; j < len(t.Attr); j++ {
// 				if t.Attr[j].Key == attribKey && t.Attr[j].Val == attribVal {
// 					return t
// 				}
// 			}
// 		}
// 		chalNode := getChallengeNode(t, atomType, tagVal)
// 		if chalNode != nil {
// 			return chalNode
// 		}
// 		t = t.NextSibling
// 	}
// 	return nil
// }
