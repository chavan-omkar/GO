package main

func main() {

	/*
		Go is not a object oriented programming language
		Try running the program by uncommenting below lines one by one and try to understand the flow
	*/

	// cards := newCard()

	// fmt.Println(card)

	// printState()

	// cards := deck{"Ace of Diamond", newCard()} //[]string{"Ace of Diamond", newCard()} //This is example of slice which is declared using square brackets

	// cards = append(cards, "Six of Spades") // This append function returns a new slice and does not modifies existing slice,used to add values in slice

	// cards := newDeck()
	// fmt.Println(cards)

	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Print("_________________******************_________________")
	// remainingCards.print()

	// greetings := "Hi there!"
	// fmt.Println([]byte(greetings)) // type conversion from string to byte slice

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	// cards := readFromFile("my_car")

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
