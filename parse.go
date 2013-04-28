package main

import (
	"exp/html"
	"fmt"
	//"strings"
	"log"
	"os"
)

func main() {
	var page *os.File
	var err error
	if page, err = os.Open("input.txt"); err != nil {
		return
	}
	defer page.Close()
	doc, err := html.Parse(page)
	//s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	//doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
