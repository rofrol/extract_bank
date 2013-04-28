package main

import (
	"code.google.com/p/go.net/html"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"rofrol/helper"
	"strconv"
)

func main() {
	var page *os.File
	var err error
	if page, err = os.Open("bank_2012.html"); err != nil {
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

		var messages []helper.Message
		for _, tr := range trArr {
			tdArr := helper.ElementsByTag(tr, "td")
			messages = append(messages, helper.String2Message(tdArr))
		}
		filename := "test.csv"

		fmt.Println("writing: " + filename)
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		writer := csv.NewWriter(f)

		headers := []string{"Opis", "Data złożenia dyspozycji", "Data waluty", "Kwota", "Saldo po operacji"}
		err = writer.Write(headers)
		if err != nil {
			log.Fatal(err)
		}

		for _, m := range messages {
			s := []string{m.Title, m.TOrd.Format("2006-01-02"), m.TExe.Format("2006-01-02"),
				strconv.FormatFloat(m.Balance, 'f', 2, 64), strconv.FormatFloat(m.Saldo, 'f', 2, 64)}
			err := writer.Write(s)
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
