package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type bookInfo struct {
	ISBN  string
	Title string
}

type bookChanel chan *bookInfo

func main() {

	filepath := "isbn-eg.csv"

	books := parseCsv(filepath)

	// var bookChan bookChanel
	bookCh := make(bookChanel)
	go func() {
		for i, book := range books {
			go func(i int, book *bookInfo) {
				booksStr := &googleapisBookJsonStruct{}
				isbn := book.ISBN
				myJsonString := googleapisBooks(isbn)
				json.Unmarshal([]byte(myJsonString), &booksStr)
				// booksStr := <-gchan
				if booksStr != nil {
					if len(booksStr.Items) > 0 {
						// fmt.Println("Find: ", isbn, " ", booksStr.Items[0].Volumeinfo.Title)
						book.Title = booksStr.Items[0].Volumeinfo.Title
					} else {
						// fmt.Println("Not found: ", isbn)
					}
				}
				bookCh <- book
			}(i, book)
		}
	}()

	bs := []*bookInfo{}
	var book *bookInfo
	for {
		book = <-bookCh
		if book == nil {
			fmt.Println("----Not Found! ----")
		}
		fmt.Print("ISBN: ", book.ISBN, " Title: ", book.Title, "\n")
		bs = append(bs, book)
	}

}

func parseCsv(filepath string) []*bookInfo {

	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var books []*bookInfo
	// records[1:] ignore header
	for _, record := range records[1:] {
		book := new(bookInfo)
		book.ISBN = record[0]
		books = append(books, book)
	}
	return books
}
