package main

import "fmt"

type deck []string

// (d deck) is called receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
