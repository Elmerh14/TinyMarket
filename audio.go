package main

import "fmt"

// enums not supported in Go but custom types can be used with const
// step 1: define a new custom type
type GenreType string

// step 2: define constants of that type
const (
	Blues     GenreType = "Blues"
	Classical GenreType = "Classical"
	Country   GenreType = "Country"
	Folk      GenreType = "Folk"
	Jazz      GenreType = "Jazz"
	Metal     GenreType = "Metal"
	Pop       GenreType = "Pop"
	RnB       GenreType = "RnB"
	Rock      GenreType = "Rock"
)

type AudioProduct struct {
	BaseProduct //embedded base struct, inherits ID, name, price, reviewRate
	Singer      NameType
	Genre       GenreType
}

// constructor for audio product
func NewAudioProduct(name string, price float64, singer NameType) *AudioProduct { //returns pointer to AudioProduct
	ap := &AudioProduct{
		BaseProduct: BaseProduct{ //initialize the base struct inside of the audio struct
			ProductID: CreateNewID(),
		},
		Singer: singer,
	}
	ap.SetProdName(name)
	ap.SetPrice(price)
	return ap
}

// getters and setters
func (ap *AudioProduct) GetSinger() NameType {
	return ap.Singer
}

func (ap *AudioProduct) SetSinger(singer NameType) {
	ap.Singer = singer
}

func (ap *AudioProduct) GetGenre() GenreType {
	return ap.Genre
}

func (ap *AudioProduct) SetGenre(genre GenreType) {
	ap.Genre = genre
}

func (ap *AudioProduct) GetProdTypeStr() string {
	return "Music"
}

// implmentation of helper functions
func (ap *AudioProduct) DisplayContentsInfo() {
	fmt.Printf("Singer Name: %s %s\n", ap.Singer.FirstName, ap.Singer.LastName)
	fmt.Printf("Genre: %s\n", ap.Genre)
}

func (ap *AudioProduct) DisplayProdInfo() {
	// Call BaseProduct display
	fmt.Printf("[Music]\n")
	fmt.Printf("Product ID: %d   Product Name: %s\n", ap.GetProdID(), ap.GetProdName())
	fmt.Printf("Price: $%.2f   Product Review Rate: %.1f\n", ap.GetPrice(), ap.GetReviewRate())
	// Then call content-specific
	ap.DisplayContentsInfo()
	fmt.Println()
}
