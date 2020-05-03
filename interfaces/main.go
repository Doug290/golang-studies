package main

import "fmt"

// if you want to fullfil/satisfy the interface bot
// you must implement a getGreeting func that returns a string
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

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generation an enlsh greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generation an enlsh greeting
	return "Hola!"
}
