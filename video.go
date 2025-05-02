package main

import "fmt"

// custom types defining
type FilmRateType string

const (
	NotRated FilmRateType = "NotRated"
	G        FilmRateType = "G"
	PG       FilmRateType = "PG"
	PG13     FilmRateType = "PG_13"
	R        FilmRateType = "R"
	NC17     FilmRateType = "NC_17"
)

type VideoProduct struct {
	BaseProduct // Embed base product fields
	Director    NameType
	FilmRate    FilmRateType
	ReleaseYear int
	RunTime     int
}

func NewVideoProduct(name string, price float64, director NameType, year, runtime int) *VideoProduct {
	vp := &VideoProduct{
		BaseProduct: BaseProduct{
			ProductID: CreateNewID(),
		},
		//composite literals set the value of variables to the value of the parameters
		Director:    director,
		ReleaseYear: year,
		RunTime:     runtime,
	}
	vp.SetProdName(name)
	vp.SetPrice(price)
	return vp
}

// getters and setters
func (vp *VideoProduct) GetDirector() NameType {
	return vp.Director
}
func (vp *VideoProduct) SetDirector(d NameType) {
	vp.Director = d
}

func (vp *VideoProduct) GetFilmRate() FilmRateType {
	return vp.FilmRate
}
func (vp *VideoProduct) SetFilmRate(rate FilmRateType) {
	vp.FilmRate = rate
}

func (vp *VideoProduct) GetReleaseYear() int {
	return vp.ReleaseYear
}
func (vp *VideoProduct) SetReleaseYear(y int) {
	vp.ReleaseYear = y
}

func (vp *VideoProduct) GetRunTime() int {
	return vp.RunTime
}
func (vp *VideoProduct) SetRunTime(r int) {
	vp.RunTime = r
}

func (vp *VideoProduct) IsNewRelease(since int) bool {
	return vp.ReleaseYear >= since
}

// implement interface methods
func (vp *VideoProduct) GetProdTypeStr() string {
	return "Movie"
}

func (vp *VideoProduct) DisplayContentsInfo() {
	fmt.Printf("Director Name: %s %s\n", vp.Director.FirstName, vp.Director.LastName)
	fmt.Printf("Film Rating: %s\n", vp.FilmRate)
	fmt.Printf("Release Year: %d\n", vp.ReleaseYear)
	fmt.Printf("Runtime: %d minutes\n", vp.RunTime)
}

func (vp *VideoProduct) DisplayProdInfo() {
	fmt.Printf("[Movie]\n")
	fmt.Printf("Product ID: %d   Product Name: %s\n", vp.GetProdID(), vp.GetProdName())
	fmt.Printf("Price: $%.2f   Product Review Rate: %.1f\n", vp.GetPrice(), vp.GetReviewRate())
	vp.DisplayContentsInfo()
	fmt.Println()
}
