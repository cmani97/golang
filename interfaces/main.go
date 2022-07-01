package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	printGreeting(eb)

	sb := spanishBot{}
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
