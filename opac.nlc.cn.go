package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func OpacDotNlcDotCn(tok, isbcode string) []byte {

	URL := "http://opac.nlc.cn/F/" + tok
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, URL, nil)

	// if you appending to existing query this works fine
	q := req.URL.Query()
	q.Add("find_base", "NLC01")
	q.Add("find_base", "NLC09")
	q.Add("func", "find-m")
	q.Add("find_code", "ISB")

	q.Add("request", isbcode)

	req.URL.RawQuery = q.Encode()

	// fmt.Println(req.URL.String())

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// fmt.Println(string(body))

	return body
}

func FoundBooksByOpacNlcCn(tok string, books []*bookInfo) (foundBooks []*bookInfo, notProcess []*bookInfo) {
	// var bookChan bookChanel
	bookCh := make(bookChanel)
	go func() {
		for i, book := range books {
			go func(i int, book *bookInfo) {
				isbn := book.ISBN
				data := OpacDotNlcDotCn(tok, isbn)
				titleraw := FindTitle(string(data))
				titleraw = strings.Replace(titleraw, "\n", "", -1)
				titleraw = strings.Trim(titleraw, " ")
				titleraw = strings.TrimRight(titleraw, " ")
				book.Title = titleraw
				bookCh <- book
			}(i, book)
		}
	}()

	var book *bookInfo
	for i := 0; i < len(books); i++ {
		book = <-bookCh
		if book == nil {
			fmt.Println("----Not Found! ----")
		}
		if book.Title == "" {
			notProcess = append(notProcess, book)
			continue
		}
		fmt.Print("ISBN: ", book.ISBN, " Title: ", book.Title, "\n")
		foundBooks = append(foundBooks, book)
	}

	return foundBooks, notProcess
}
