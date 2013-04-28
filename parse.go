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
	// bytes.Buffer does not implement io.Writer
	// output needs to be pointer 
	// http://my.safaribooksonline.com/book/programming/9780132918961/a-go-primer/ch02lev1sec8
	// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/xXFpT8oLGNU
	// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/jDcbmzDaLi8
	//var output bytes.Buffer
	//than html.Render(&ouput, node)
	//or
	//output := new(bytes.Buffer)
	//than html.Render(output, node)
	var page *os.File
	var err error
	if page, err = os.Open("bank_2012.html"); err != nil {
		panic(err)
	}
	defer page.Close()

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
			// https://groups.google.com/d/msg/golang-nuts/ybwJH4pR0lo/gw38xJgY3AoJ
			// "%.3f" rounds at the third place after the suffix and truncates to two places after the suffix.
			//If the suffix would be all zero, it is completely discarded. 
			//0.5 is rounded towards zero. 
			s := []string{m.Title, m.TOrd.Format("2006-01-02"), m.TExe.Format("2006-01-02"),
				strconv.FormatFloat(m.Balance, 'f', 2, 64), strconv.FormatFloat(m.Saldo, 'f', 2, 64)}
			err := writer.Write(s)
			if err != nil {
				log.Fatal(err)
			}
		}
		f.Close()
	}
}
