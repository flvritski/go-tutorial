package main

func main() {

	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")

	//cards2 := newDeckFromFile("my_cards")
	//cards2.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
