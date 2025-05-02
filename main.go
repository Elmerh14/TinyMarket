package main

func main() {
	//Audion Products
	audio1 := NewAudioProduct("Yesterday", 16.5, NameType{"Beatles", ""})
	audio1.SetGenre(Pop)
	audio1.SetReviewRate(9.8)

	audio2 := NewAudioProduct("We Are the World", 13.75, NameType{"Michael", "Jackson"})
	audio2.SetGenre(Country)
	audio2.SetReviewRate(9.1)

	audio3 := NewAudioProduct("Take Me Home", 12.50, NameType{"John", "Denver"})
	audio3.SetGenre(Folk)
	audio3.SetReviewRate(9.9)

	//Video Products
	video1 := NewVideoProduct("Sound of Music", 22.0, NameType{"Robert", "Wise"}, 1965, 175)
	video1.SetFilmRate(G)
	video1.SetReviewRate(9.2)

	video2 := NewVideoProduct("Star Wars", 22.0, NameType{"George", "Lucas"}, 1977, 120)
	video2.SetFilmRate(PG)
	video2.SetReviewRate(8.5)

	//Books
	ebook := NewEBook("The Old Man and the Sea", 8.3, NameType{"Ernest", "Hemingway"}, 127)
	ebook.SetReviewRate(9.5)

	paperbook := NewPaperBook("Grapes of Wrath", 10.5, NameType{"John", "Steinbeck"}, 150)
	paperbook.SetReviewRate(9.0)

	//create the cart and add products

	cart := NewCart(NameType{"Jhon", "Smith"})

	cart.AddItem(audio1)
	cart.AddItem(audio2)
	cart.AddItem(audio3)
	cart.AddItem(video1)
	cart.AddItem(video2)
	cart.AddItem(ebook)
	cart.AddItem(paperbook)

	//remove two products from the cart
	cart.RemoveItem(audio2.GetProdID())
	cart.RemoveItem(paperbook.GetProdID())

	//disply cart contents
	cart.DisplayCart()

}
