package main

func main() {
	cards := newDeck()
	cards.print()

	remaining, hand := deal(cards, 5)
	remaining.print()
	hand.print()
}
