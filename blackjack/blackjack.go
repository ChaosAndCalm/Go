package blackjack

import "strings"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	value := -1

	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "queen":
		value = 10
	case "king":
		value = 10
	case "joker":
		value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1 string, card2 string, dealerCard string) string {

	var val1 int = ParseCard(card1)
	var val2 int = ParseCard(card2)

	var dealerVal int = ParseCard(dealerCard)

	if strings.Compare(card1, "ace") == 0 && strings.Compare(card2, "ace") == 0 {
		return "P"
	}

	if val1+val2 == 21 {
		if dealerVal < 10 {
			return "W"
		} else {
			return "S"
		}
	}

	if val1+val2 >= 17 && val1+val2 <= 20 {
		return "S"
	}

	if val1+val2 >= 12 && val1+val2 <= 16 {
		if dealerVal < 7 {
			return "S"
		} else {
			return "H"
		}
	}

	return "H"
}
