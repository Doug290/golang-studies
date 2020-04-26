package main

func main() {
	cards := newDeck()
	cards.saveToFile("myCards")
}