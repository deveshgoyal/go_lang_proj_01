package main

/*
import(
	"deck"
)
*/

func main() {
	d := Deck{}
	// d.reset()
	// deal := d.deal(4)
	d.reset()
	d.shuffle()
	// d.save("first.txt")
	// d.load("first.txt")
	d.print()
	// fmt.Println("deck:")
	// d.print()
	// fmt.Println("deal cards:")
	// deal.save("first.txt")
	d.print()
}
