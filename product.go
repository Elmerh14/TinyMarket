package main

// name type struct
type NameType struct {
	FirstName string
	LastName  string
}

// Interface for all products — simulates abstraction and polymorphism capital letter means interface will be public
// accessible to all outside code aslong as package is imported
type Product interface {
	GetProdID() int
	GetProdName() string
	GetPrice() float64
	GetReviewRate() float32
	SetProdID(int)
	SetProdName(string)
	SetPrice(float64)
	SetReviewRate(float32)

	GetProdTypeStr() string
	DisplayProdInfo()
	DisplayContentsInfo()
}

// base struct to be embedded in the other structs. Think of this as like the abstract class in c++
type BaseProduct struct {
	ProductID   int
	ProductName string
	Price       float64
	ReviewRate  float32
}

// shared ID generator (simulating static variable)
var nextID int = 1

func CreateNewID() int {
	id := nextID
	nextID++
	return id
}

// method definitions - structs embedding BaseProduct will inherit these methods

func (bp *BaseProduct) GetProdID() int         { return bp.ProductID } //*BaseProduct means pointer reciver, bp pointer to BaseProduct struct not struct itself.
func (bp *BaseProduct) GetProdName() string    { return bp.ProductName }
func (bp *BaseProduct) GetPrice() float64      { return bp.Price }
func (bp *BaseProduct) GetReviewRate() float32 { return bp.ReviewRate }
func (bp *BaseProduct) SetProdID(id int)       { bp.ProductID = id }

func (bp *BaseProduct) SetProdName(name string) {
	if name == "" {
		bp.ProductName = "!No Name Product!"
	} else {
		bp.ProductName = name
	}
}

func (bp *BaseProduct) SetPrice(price float64) {
	if price > 0 && price < 1000 {
		bp.Price = price
	}
}

func (bp *BaseProduct) SetReviewRate(rate float32) {
	bp.ReviewRate = rate
}
