package money

import (
	"errors"
	"fmt"

	"github.com/Mr-MichaelK/money-man/internal/core/currency"
)

var (
	ErrNilCurrencyFetcher  = errors.New("currency fetcher cannot be nil")
	ErrInvalidCurrencyCode = errors.New("invalid currency code")
)

// CurrencyFetcher takes a CurrencyCode and returns its corresponding currency.
type CurrencyFetcher interface {
	Fetch(currencyCode currency.CurrencyCode) (currency.Currency, error)
}

// MoneyFactory is the standard way to create Money with a certain currency. Prefer using it over
// the money.New constructor to avoid having to create a new currency and to validate the one being
// used is valid and exists.
type MoneyFactory struct {
	currencyFetcher CurrencyFetcher
}

// NewMoneyFactory takes a CurrencyFetcher and returns a new money factory.
func NewMoneyFactory(currencyFetcher CurrencyFetcher) (*MoneyFactory, error) {
	if currencyFetcher == nil {
		return &MoneyFactory{}, ErrNilCurrencyFetcher
	}

	return &MoneyFactory{currencyFetcher: currencyFetcher}, nil
}

// NewMoney creates money from a certain amount and a CurrencyCode.
func (m *MoneyFactory) NewMoney(amount Amount, currencyCode currency.CurrencyCode) (Money, error) {
	if !currencyCode.IsValid() {
		return Money{}, fmt.Errorf("%w: %q", ErrInvalidCurrencyCode, currencyCode)
	}

	currency, err := m.currencyFetcher.Fetch(currencyCode)
	if err != nil {
		return Money{}, fmt.Errorf("an error occurred when fetching currency %s: %w", currencyCode, err)
	}

	return New(amount, currency), nil
}
