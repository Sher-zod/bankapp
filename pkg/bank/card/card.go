package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}
func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Active == true {
		if amount < card.Balance {
			if amount > 0 {
				if amount < 20_000_00 {
					card.Balance = card.Balance - amount
				}
			}
		}
	} else {
		card.Balance = card.Balance * 1
	}
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	if amount < 0 {
		return
	}

	if amount > 50_000_00 {
		return
	}

	if !(*card).Active {
		return
	}

	(*card).Balance += amount
}

const bonusLimit = 5_000_00

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !(*card).Active {
		return
	}
	if card.Balance <= 0 {
		return
	}

	bonus := int(card.MinBalance) * percent / 100 * daysInMonth / daysInYear

	if bonus > bonusLimit {
		return
	}

	(*card).Balance += types.Money(bonus)

}

func Total(cards []types.Card) types.Money {
	sumBalance := types.Money(0)
	for _, card := range cards {
		if card.Active && card.Balance >= 0 {
			sumBalance += card.Balance
		}

	}
	return sumBalance
}