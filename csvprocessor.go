package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func booksWriter(csvfilename string, books []*bookInfo) {
	// var books []*bookInfo
	// 1. Open the file
	recordFile, err := os.Create("./" + csvfilename)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	var csvData = [][]string{{"isbn", "title"}}

	// 2. Initialize the writer
	writer := csv.NewWriter(recordFile)
	for _, b := range books {
		binfo := []string{b.ISBN, b.Title}
		csvData = append(csvData, binfo)
	}

	// 3. Write all the records
	err = writer.WriteAll(csvData) // returns error
	if err != nil {
		fmt.Println("An error encountered ::", err)
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
