package main

func main() {
	// cards := newDeck()
	// cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// fmt.Println("string to byte convertion")
	// fmt.Println([]byte("Ht there!"))

	// stringOfCards := cards.toString()
	// fmt.Println("stringOfCards= ", stringOfCards)

	// fmt.Println("saving the cards")
	// cards.saveToFile("MyCards")

	// fmt.Println("Reading cards from file")
	// newCardsFromFile := newDeckFromFile("MyCards1")
	// newCardsFromFile.print()

	cards := newDeck()
	cards.suffle()
	cards.print()
}

// type laptopSize float64
// func (this laptopSize) getSizeOfLaptop() laptopSize {
// 	return this
// }

// package main
// import "fmt"
// var deckSize int
// func main() {
// 	deckSize = 50
// 	fmt.Println(deckSize)
// }
