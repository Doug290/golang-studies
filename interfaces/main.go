package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting (eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generation an enlsh greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generation an enlsh greeting
	return "Hola!"
}
