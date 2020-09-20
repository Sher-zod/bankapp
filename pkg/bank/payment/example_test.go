package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 20_000_00,
		},
		{
			ID:     2,
			Amount: 30_000_00,
		},
		{
			ID:     3,
			Amount: 50_000_00,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 3

}
