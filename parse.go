package main

import (
	"exp/html"
	"fmt"
	"log"
	"os"
	"rofrol/helper"
	"bytes"
)

var table *html.Node
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
					if t {
						table = n
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	child = div.FirstChild
	html.Render(output, child)

/*
	for child = child.NextSibling; child == nil; child = child.NextSibling {}
	html.Render(output, child)

	for child = child.NextSibling; child == nil; child = child.NextSibling {}
	html.Render(output, child)
*/

	accounts_history_1 := string(output.Bytes())
	output.Reset()

	td = td.NextSibling
	html.Render(output, td.FirstChild)
	accounts_history_2 := string(output.Bytes())
	output.Reset()

	td = td.NextSibling
	html.Render(output, td.FirstChild)
	accounts_history_3 := string(output.Bytes())
	output.Reset()

	td = td.NextSibling
	html.Render(output, td.FirstChild)
	accounts_history_4 := string(output.Bytes())
	output.Reset()

	td = td.NextSibling
	html.Render(output, td.FirstChild)
	accounts_history_5 := string(output.Bytes())
	output.Reset()

	fmt.Println("1", accounts_history_1)
	fmt.Println("2", accounts_history_2)
	fmt.Println("3", accounts_history_3)
	fmt.Println("4", accounts_history_4)
	fmt.Println("5", accounts_history_5)
}
