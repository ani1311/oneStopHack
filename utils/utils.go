package utils

import (
	"bytes"
	// "fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"io/ioutil"
	"net/http"
)

func GetPage(page string) io.ReadCloser {
	resp, _ := http.Get(page)
	return resp.Body
}

func GetPageFromLocal(page string) io.ReadCloser {
	dat, _ := ioutil.ReadFile(page)
	return ioutil.NopCloser(bytes.NewReader(dat))
}

func GetChallengeNodeUsingAttrib(n *html.Node, atomType atom.Atom, attribKey string, attribVal string) *html.Node {
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
		chalNode := GetChallengeNodeUsingAttrib(t, atomType, attribKey, attribVal)
		if chalNode != nil {
			return chalNode
		}
		t = t.NextSibling
	}
	return nil
}

func GetChallengeNodeUsingAtom(n *html.Node, atomType atom.Atom, tagVal string) *html.Node {
	if n == nil {
		return nil
	}
	t := n.FirstChild
	for t != nil {
		if t.DataAtom == atomType && t.FirstChild.Data == tagVal {
			return t
		}
		chalNode := GetChallengeNodeUsingAtom(t, atomType, tagVal)
		if chalNode != nil {
			return chalNode
		}
		t = t.NextSibling
	}
	return nil
}
