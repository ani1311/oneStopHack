package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func getPage(page string) io.ReadCloser {
	resp, _ := http.Get(page)
	return resp.Body
}

func getPageFromLocal(page string) io.ReadCloser {
	dat, _ := ioutil.ReadFile(page)
	return ioutil.NopCloser(bytes.NewReader(dat))
}
