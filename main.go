package main

type bookInfo struct {
	ISBN  string
	Title string
}

type bookChanel chan *bookInfo

func main() {
	filepath := "isbn-eg.csv"
	books := parseCsv(filepath)
	// FindByGoogleWithISBN(books)
	FindByOpacNlcCnWithISBN(books)
}

func FindByGoogleWithISBN(books []*bookInfo) {
	foundBooks, notProcess := FoundBooksByGooglleapis(books)
	booksWriter("found-title-by-google.csv", foundBooks)
	booksWriter("not-found-by-google.csv", notProcess)
}
func FindByOpacNlcCnWithISBN(books []*bookInfo) {
	tok := "G6FSUJ9VTBXXQ6YBFPPFN6DHIPUGNFHHMAN3RY7T7EGG77C7JT-17240"
	foundBooks, notProcess := FoundBooksByOpacNlcCn(tok, books)
	booksWriter("found-title-by-opacnlc.csv", foundBooks)
	booksWriter("not-found-by-opacnlc.csv", notProcess)
}
