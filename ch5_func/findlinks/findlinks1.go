package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(target string) io.Reader {
	res, err := http.Get(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reding %s: %v\n", target, err)
		os.Exit(1)
	}
	return bytes.NewReader(b)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func findLink(target string) {
	doc, err := html.Parse(fetch(target))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func main() {
	findLink("https://baidu.com")
}
