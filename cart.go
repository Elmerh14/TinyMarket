package main

import "fmt"

const MAX_ITEMS int = 7

type Cart struct {
	ItemNum int
	Owner   NameType
	//A slice of the interface type, allows polymorphis any struct that fully implements the product interface a pointer to the actual value will be stored here
	PurchasedItems []Product
}

// constructor for cart
func NewCart(owner NameType) *Cart {
	return &Cart{
		Owner:          owner,
		PurchasedItems: make([]Product, 0, MAX_ITEMS), // pre-allocate space
	}
}

// helper functions

func (c *Cart) isCartFull() bool {
	return len(c.PurchasedItems) >= MAX_ITEMS
}

// append purchased items to the slice if the cart is not full
func (c *Cart) AddItem(p Product) bool {
	if c.isCartFull() {
		fmt.Println("Cart is full. Cannot add more items.")
		return false
	}
	c.PurchasedItems = append(c.PurchasedItems, p)
	c.ItemNum++
	return true
}

func (c *Cart) RemoveItem(productID int) bool {
	for i, item := range c.PurchasedItems {
		if item.GetProdID() == productID {
			//slice deletion pattern, all all elements before index i, all elements after index i, gule them together
			c.PurchasedItems = append(c.PurchasedItems[:i], c.PurchasedItems[i+1:]...)
			c.ItemNum--
			return true
		}
	}
	fmt.Println("Product not found in cart.")
	return false
}

func (c *Cart) DisplayCart() {
	fmt.Println("My Cart")
	fmt.Println("======")
	fmt.Printf("Cart Owner: %s %s\n\n", c.Owner.FirstName, c.Owner.LastName)

	var total float64 = 0.0

	for _, item := range c.PurchasedItems {
		//go will know which implemented interface to call, audioProduct, videProduct, eBookProduct...
		item.DisplayProdInfo()
		total += item.GetPrice()
	}

	count := len(c.PurchasedItems)
	avg := 0.0
	if count > 0 {
		avg = total / float64(count)
	}

	fmt.Println("===== Summary of Purchase ======")
	fmt.Printf("Total number of purchases: %d\n", count)
	fmt.Printf("Total purchasing amount: $%.2f\n", total)
	fmt.Printf("Average cost: $%.2f\n", avg)
}
