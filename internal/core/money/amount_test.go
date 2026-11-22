package money_test

import (
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/money"
)

func TestAmountAdd(t *testing.T) {
	const (
		startingAmount = 0
		addedAmount    = 100
	)
	t.Run("given an amount, the system should add it to itself in place", func(t *testing.T) {
		baseAmount := money.Amount(startingAmount)

		baseAmount.Add(money.Amount(addedAmount))

		expectedAmount := money.Amount(startingAmount + addedAmount)
		if baseAmount != expectedAmount {
			t.Errorf("expected amount to become %v but got %v", expectedAmount, baseAmount)
		}
	})
}

func TestAmountSubtract(t *testing.T) {
	const (
		startingAmount   = 100
		subtractedAmount = 20
	)
	t.Run("given an amount, the system should subtract it from itself in place", func(t *testing.T) {
		baseAmount := money.Amount(startingAmount)

		baseAmount.Subtract(money.Amount(subtractedAmount))

		expectedAmount := money.Amount(startingAmount - subtractedAmount)
		if baseAmount != expectedAmount {
			t.Errorf("expected amount to become %v but got %v", expectedAmount, baseAmount)
		}
	})
}
