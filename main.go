package main
type deck []string

func newDeck() deck {
	cards := deck{}

	suits := deck{"1", "2", "3", "4"}
	values := deck{"a", "b", "c"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, hanSize int) (deck, deck) {
	return d[:hanSize], d[hanSize:]
}

func main() {
	cards := newDeck()

	hand, remain := deal(cards, 2)
	hand.print()
	remain.print()
}
