package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func googleapisBooks(isbcode string) []byte {

	URL := `https://www.googleapis.com/books/v1/volumes?q=isbn:` + isbcode
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, URL, nil)

	fmt.Println(req.URL.String())

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

func FoundBooksByGooglleapis(books []*bookInfo) (foundBooks []*bookInfo, notProcess []*bookInfo) {
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

type googleapisBookJsonStruct struct {
	Kind       string `json:"kind"`
	Totalitems int    `json:"totalItems"`
	Items      []struct {
		Kind       string `json:"kind"`
		ID         string `json:"id"`
		Etag       string `json:"etag"`
		Selflink   string `json:"selfLink"`
		Volumeinfo struct {
			Title               string `json:"title"`
			Subtitle            string `json:"subtitle"`
			Publisheddate       string `json:"publishedDate"`
			Industryidentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
			Readingmodes struct {
				Text  bool `json:"text"`
				Image bool `json:"image"`
			} `json:"readingModes"`
			Printtype           string `json:"printType"`
			Maturityrating      string `json:"maturityRating"`
			Allowanonlogging    bool   `json:"allowAnonLogging"`
			Contentversion      string `json:"contentVersion"`
			Panelizationsummary struct {
				Containsepubbubbles  bool `json:"containsEpubBubbles"`
				Containsimagebubbles bool `json:"containsImageBubbles"`
			} `json:"panelizationSummary"`
			Language            string `json:"language"`
			Previewlink         string `json:"previewLink"`
			Infolink            string `json:"infoLink"`
			Canonicalvolumelink string `json:"canonicalVolumeLink"`
		} `json:"volumeInfo,omitempty"`
		Saleinfo struct {
			Country     string `json:"country"`
			Saleability string `json:"saleability"`
			Isebook     bool   `json:"isEbook"`
		} `json:"saleInfo"`
		Accessinfo struct {
			Country                string `json:"country"`
			Viewability            string `json:"viewability"`
			Embeddable             bool   `json:"embeddable"`
			Publicdomain           bool   `json:"publicDomain"`
			Texttospeechpermission string `json:"textToSpeechPermission"`
			Epub                   struct {
				Isavailable bool `json:"isAvailable"`
			} `json:"epub"`
			Pdf struct {
				Isavailable bool `json:"isAvailable"`
			} `json:"pdf"`
			Webreaderlink       string `json:"webReaderLink"`
			Accessviewstatus    string `json:"accessViewStatus"`
			Quotesharingallowed bool   `json:"quoteSharingAllowed"`
		} `json:"accessInfo"`
		Searchinfo struct {
			Textsnippet string `json:"textSnippet"`
		} `json:"searchInfo,omitempty"`
	} `json:"items"`
}
