package main

import "fmt"

// base struct for book
type BookProduct struct {
	BaseProduct
	Author NameType
	Pages  int
}

// define a constructor for book, book wont be instantiated but constructor will be used by e-book and paper book.
func NewBookProduct(name string, price float64, author NameType, pages int) BookProduct {
	bp := BookProduct{
		BaseProduct: BaseProduct{
			ProductID: CreateNewID(),
		},
		Author: author,
		Pages:  pages,
	}
	bp.SetProdName(name)
	bp.SetPrice(price)
	return bp //returns an actual struct that will be embedded inside ebook and paper book
}

// book getters and setters
func (bp *BookProduct) GetAuthor() NameType {
	return bp.Author
}

func (bp *BookProduct) SetAuthor(a NameType) {
	bp.Author = a
}

func (bp *BookProduct) GetPages() int {
	return bp.Pages
}

func (bp *BookProduct) SetPages(p int) {
	bp.Pages = p
}

//define Ebook and PaperBook

type EBook struct {
	BookProduct
}

type PaperBook struct {
	BookProduct
}

// consturctors for Ebook and PaperBook
func NewEBook(name string, price float64, author NameType, pages int) *EBook {
	return &EBook{
		BookProduct: NewBookProduct(name, price, author, pages),
	}
}

func NewPaperBook(name string, price float64, author NameType, pages int) *PaperBook {
	return &PaperBook{
		BookProduct: NewBookProduct(name, price, author, pages),
	}
}

//both EBook and PaperBook myst implement the Product interface

// EBook interface implementation
func (eb *EBook) GetProdTypeStr() string {
	return "E book"
}

func (eb *EBook) DisplayContentsInfo() {
	fmt.Printf("Author: %s %s\n", eb.Author.FirstName, eb.Author.LastName)
	fmt.Printf("Pages: %d\n", eb.Pages)
}

func (eb *EBook) DisplayProdInfo() {
	fmt.Printf("[E book]\n")
	fmt.Printf("Product ID: %d   Product Name: %s\n", eb.GetProdID(), eb.GetProdName())
	fmt.Printf("Price: $%.2f   Product Review Rate: %.1f\n", eb.GetPrice(), eb.GetReviewRate())
	eb.DisplayContentsInfo()
	fmt.Println()
}

// PaperBook interface implementation
func (pb *PaperBook) GetProdTypeStr() string {
	return "Paper book"
}

func (pb *PaperBook) DisplayContentsInfo() {
	fmt.Printf("Author: %s %s\n", pb.Author.FirstName, pb.Author.LastName)
	fmt.Printf("Pages: %d\n", pb.Pages)
}

func (pb *PaperBook) DisplayProdInfo() {
	fmt.Printf("[Paper book]\n")
	fmt.Printf("Product ID: %d   Product Name: %s\n", pb.GetProdID(), pb.GetProdName())
	fmt.Printf("Price: $%.2f   Product Review Rate: %.1f\n", pb.GetPrice(), pb.GetReviewRate())
	pb.DisplayContentsInfo()
	fmt.Println()
}
