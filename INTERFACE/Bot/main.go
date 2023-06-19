package main

import "fmt"

/*
	Due to this declaration the other types that implement function

of name getGreeting and returns string makes the other types honorory memeber of type bot
*/
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

/*
functions getGreeting are expecting receiver's of type
englishBot and spanishBot
*/

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*

Types we are using here are

1) Concrete Type :- map struct,int string,englishBot
we can create value directly from the above mentioned so those are called concrete

2) Interface Type :- bot
we cannot directly create value from this type bot so we called it as interface type

*/
