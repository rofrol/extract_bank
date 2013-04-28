package main

import (
	"exp/html"
	"fmt"
	"log"
	"os"
	"rofrol/helper"
)

func main() {
	var page *os.File
	var err error
	if page, err = os.Open("bank_2012.html"); err != nil {
		return
	}
	defer page.Close()

	doc, err := html.Parse(page)
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			for _, a := range n.Attr {
				if a.Key == "class" {
					fmt.Println(a.Val)
					t := helper.FoundClass(a.Val, "contentRoles6")
					fmt.Println("found class?", t)
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
