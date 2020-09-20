package main

import (
	"bank/pkg/bank/card"
	"fmt"
)

func main() {
	card := card.IssueCard("TJS", "green", "extreme")
	fmt.Printf("%+v", card)
	
}
