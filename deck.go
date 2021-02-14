package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, numberOfCards int) (deck, deck) {
	return d[:numberOfCards], append(d[numberOfCards:])
}

func createFromFile(filename string) (deck, error) {

	text, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	s := string(text)

	d := fromString(s)

	return d, nil
}

func fromString(s string) deck {
	spl := strings.Split(s, ", ")
	d := deck(spl)

	return d
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) saveToFile(filename string) error {
	deckAsBytes := []byte(d.toString())

	err := ioutil.WriteFile(filename, deckAsBytes, os.ModePerm)

	if err != nil {
		return nil
	}

	return err
}

func (d deck) toString() string {
	stringSlice := []string(d)

	asString := strings.Join(stringSlice, ", ")

	return asString
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
