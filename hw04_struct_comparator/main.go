package main

import "fmt"

type BookField int

const (
	Rate BookField = iota // on demand == 0
	Year
	Size
)

type BookCompare struct {
	bookField BookField
}

func NewBookCompare(bf BookField) BookCompare {
	return BookCompare{
		bookField: bf,
	}
}

func (bc BookCompare) CompareBooks(b1, b2 Book) bool {
	switch bc.bookField {
	case Rate:
		return b1.rate == b2.rate
	case Year:
		return b1.year == b2.year
	case Size:
		return b1.size == b2.size
	}

	return true
}

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func SetID(id int, b *Book) {
	b.id = id
}

func SetTitle(title string, b *Book) {
	b.title = title
}

func SetAuthor(author string, b *Book) {
	b.author = author
}

func SetYear(year int, b *Book) {
	b.year = year
}

func SetSize(size int, b *Book) {
	b.size = size
}

func SetRate(rate float32, b *Book) {
	b.rate = rate
}

func GetID(b Book) int {
	return b.id
}

func GetTitle(b Book) string {
	return b.title
}

func GetAuthor(b Book) string {
	return b.author
}

func GetYear(b Book) int {
	return b.year
}

func GetSize(b Book) int {
	return b.size
}

func GetRate(b Book) float32 {
	return b.rate
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Place your code here.
	bookID := intSeq()
	book1 := Book{
		id:     bookID(),
		title:  "",
		author: "",
		year:   1981,
		size:   422,
		rate:   18.23,
	}
	book2 := Book{
		id:     bookID(),
		title:  "",
		author: "",
		year:   1981,
		size:   214,
		rate:   22.18,
	}

	bc := BookCompare{Rate}
	fmt.Println(bc.CompareBooks(book1, book2))
	bc.bookField = Year
	fmt.Println(bc.CompareBooks(book1, book2))
	bc.bookField = Size
	fmt.Println(bc.CompareBooks(book1, book2))
}
