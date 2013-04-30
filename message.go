package main

import (
	"code.google.com/p/go.net/html"
	"github.com/rofrol/helper"
	"log"
	"strconv"
	"strings"
	"time"
)

// getters setters http://stackoverflow.com/questions/11810218/how-to-set-and-get-fields-in-golang-structs
type Message struct {
	Title   string
	TOrd    time.Time
	TExe    time.Time
	Balance float64
	Saldo   float64
}

func (m Message) ArrString() []string {
	return []string{m.Title, m.TOrd.Format("2006-01-02"), m.TExe.Format("2006-01-02"),
		strconv.FormatFloat(m.Balance, 'f', 2, 64), strconv.FormatFloat(m.Saldo, 'f', 2, 64)}
}

func Headers() []string {
	return []string{"Opis", "Data złożenia dyspozycji", "Data waluty", "Kwota", "Saldo po operacji"}
}

func String2Message(arr []*html.Node) Message {
	// https://sites.google.com/site/gopatterns/error-handling
	var err error
	title := helper.String2CsvCell(arr[2].FirstChild.FirstChild.Data)
	tOrd, err := time.Parse("02.01.2006", arr[3].FirstChild.Data)
	if err != nil {
		log.Fatal(err)
	}
	tExe, err := time.Parse("02.01.2006", arr[4].FirstChild.Data)
	if err != nil {
		log.Fatal(err)
	}
	balance, err := strconv.ParseFloat(strings.Replace(arr[5].FirstChild.Data, " ", "", -1), 64)
	if err != nil {
		log.Fatal(err)
	}
	saldo, err := strconv.ParseFloat(strings.Replace(arr[6].FirstChild.Data, " ", "", -1), 64)
	if err != nil {
		log.Fatal(err)
	}

	return Message{title, tOrd, tExe, balance, saldo}
}
