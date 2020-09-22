package stats

import (
	"github.com/MrSukhrob96/bank/pkg/bank/types"
)

func Avg(pyments []types.Payment) types.Money {
	var summ types.Money = 0
	for _, pyment := range pyments {
		if pyment.Amount > 0 {
			summ = summ + pyment.Amount
		}
	}

	avg := summ / types.Money(len(pyments))
	return avg
}
