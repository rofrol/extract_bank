package main

import (
	"code.google.com/p/go.net/html"
	"encoding/csv"
	"fmt"
	"github.com/rofrol/helper"
	"log"
	"os"
)

func Rows(filename string) {
	var page *os.File
	var err error
	if page, err = os.Open(filename); err != nil {
		panic(err)
	}

	defer func() {
		if err := page.Close(); err != nil {
			panic(err)
		}
	}()

	doc, err := html.Parse(page)
	if err != nil {
		log.Fatal(err)
	}

	table := helper.FindByClass(doc, "table", "content")
	if table != nil {
		tbody := helper.FirstChildByTag(table, "tbody")
		trArr := helper.ElementsByTag(tbody, "tr")

		var messages []Message
		for _, tr := range trArr {
			tdArr := helper.ElementsByTag(tr, "td")
			messages = append(messages, String2Message(tdArr))
		}
		filename := "test.csv"

		fmt.Println("writing: " + filename)
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		writer := csv.NewWriter(f)

		err = writer.Write(Headers())
		if err != nil {
			log.Fatal(err)
		}

		for _, m := range messages {
			err := writer.Write(m.ArrString())
			if err != nil {
				log.Fatal(err)
			}
		}

		defer func() {
			if err := f.Close(); err != nil {
				panic(err)
			}
		}()

		writer.Flush()
	}
}

func main() {
	Rows(os.Args[1])
}
