package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FindTitle(data string) (title string) {

	defer func() {
		if e := recover(); e != nil {
			// log.Println(e)
			return
		}
	}()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	foundIndex := false

	// Find each table
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				if foundIndex {
					if title == "" {
						title = tablecell.Text()
						panic("Find title: " + title)
						// panic and return
					}
				}
				if strings.EqualFold(strings.TrimSpace(tablecell.Text()), "题名与责任") {
					foundIndex = true
					// log.Println("Found!")
				}
			})
		})
	})
	return title
}
