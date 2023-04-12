package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
We will create  new type of 'deck' which is a slice of string
*/

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" Of "+suits)
		}
	}

	return cards

}

/*
Any variable with type deck can access this print function
or we can say this is a function which accepts type deck as receiver
*/
func (d deck) print() {
	/*
		below is a for loop in this we are redeclaring card variable every time so we are using ':=' symbol
	*/
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
Multiple return values function example
here we have passed deck as argument and not as reciver because it would create a ambiguity
*/

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] //starting colon indicate that slice is from starting to handsize in the deck, end colon indicate that slice start from handsize to end of the deck returning two separate values
}

// returns a single string with comma seperation
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/*
io operations like saving data in a file on machine harddrive
The receiver is the object on what you declare your method [here it is d deck]
When want to add a method to an object, you use this syntax.
*/
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

/*
io operations like reading data from a file on machine harddrive
The receiver is the object on what you declare your method [here it is d deck]
When want to add a method to an object, you use this syntax.
*/

func readFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)

	// following code blcok shows error handling in go
	if err != nil {
		//Option #1 - log the error and return a call to newDeck()
		//Option #2 - log the error and entirely quit the program

		fmt.Println("Error :", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffle the cards every time the function is called
// it desn't require any other arg and all can be done using receiver and it doesnt return anything
func (d deck) shuffle() {
	for i := range d {
		// random number generation between 0 to n but here it will generate same set every time as it will use the same satrting point lets say seed in Go
		// newPosition := rand.Intn(len(d) - 1)
		//d[i], d[newPosition] = d[newPosition], d[i]

		//time.Now().UnixNano() we are using as seed to generate source using rand.NewSource() and that new source is being used to generate new random numbers everytime
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
