package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

type Deck []Card

var special = map[string]int{
	"ace":   1,
	"jack":  2,
	"queen": 3,
	"king":  4,
}

var special_re = map[int]string{
	1: "ace",
	2: "jack",
	3: "queen",
	4: "king",
}

var colors []string = []string{"club", "diamonds", "hearts", "spades"}

// TODO : Check in colors for invali]]]]]]]]d colors
// func (s colors) indexOf() {

// }

func (d *Deck) reset() {
	for c := 0; c < len(colors); c++ {
		for i := 1; i <= 13; i++ {
			// index := 13*c + i
			ca := Card{
				color: colors[c],
				value: i,
			}
			*d = append(*d, ca)
		}
	}
}

func (d *Deck) deal(n int) Deck {
	da := *d
	decks := da[:n]
	*d = da[n:]
	return decks
}

func (d *Deck) print() {
	for _, c := range *d {
		c.print()
	}
}

func (d *Deck) load(path string) {
	b, _ := ioutil.ReadFile(path)
	s := strings.Split(string(b), "\n")
	for _, cs := range s {

		i := strings.Index(cs, "Card is ")
		j := strings.Index(cs, " of ")

		value, err := strconv.Atoi(cs[i+8 : j])
		if err != nil {
			value = special[cs[i+8:j]]
		}
		if value == 0 { //|| colors  {
			continue
		}
		c := Card{
			color: cs[j+4:],
			value: value,
		}
		*d = append(*d, c)
	}

}

func (d Deck) save(path string) {
	s := []string{}
	for _, c := range d {
		s = append(s, c.string())
	}
	fmt.Println(s)
	ioutil.WriteFile(path, []byte(strings.Join(s, "\n")), 0666)
}

func (d *Deck) shuffle() {
	for i, c := range *d {
		r := rand.Intn(len(*d) - 1)
		(*d)[i], (*d)[r] = (*d)[r], c
	}
}

type Card struct {
	color string
	value int
}

func (c *Card) stringValue() string {
	if c.value == 1 {
		return "ace"
	} else if c.value == 11 {
		return "jack"
	} else if c.value == 12 {
		return "queen"
	} else if c.value == 13 {
		return "king"
	} else {
		return strconv.Itoa(c.value)
	}
}

func (c Card) string() string {
	return "Card is " + c.stringValue() + " of " + c.color
}

func (c *Card) print() {
	fmt.Println(c.string())
}
