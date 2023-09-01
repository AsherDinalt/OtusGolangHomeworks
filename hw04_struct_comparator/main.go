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

func NewBookCompare(bf BookField) *BookCompare {
	return &BookCompare{
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

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func GetSize(b Book) int {
	return b.size
}

func (b *Book) Rate() float32 {
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
