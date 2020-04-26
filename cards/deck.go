package main

import "fmt"

// Create a new type of "deck"
// which is a slice of strings
type deck []string

// d here are just the same as "this" in js or "self" in pyton
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
