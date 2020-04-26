package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("myCards")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}