package main

// import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = newCard()
	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	// cards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }

// Go is statically typed language