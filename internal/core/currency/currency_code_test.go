package currency_test

import (
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/currency"
)

func TestIsValidCurrencyCode(t *testing.T) {
	t.Run("should return false if the input is not 3 capitalized letters", func(t *testing.T) {
		invalidCurrencyCode := currency.CurrencyCode("2nf9")

		result := invalidCurrencyCode.IsValid()

		if result {
			t.Fatalf("expected IsValid to return false with currency code %v but it returned true", invalidCurrencyCode)
		}
	})

	t.Run("should return true if the input is 3 capitalized letters", func(t *testing.T) {
		validCurrencyCode := currency.CurrencyCode("USD")

		result := validCurrencyCode.IsValid()

		if !result {
			t.Fatalf("expected IsValid to return true with currency code %v but it returned false", validCurrencyCode)
		}
	})
}
