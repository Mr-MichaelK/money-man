package currency

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidCurrencyCode = errors.New("currency code should be 3 capital letters")
	ErrEmptyCurrencyName   = errors.New("currency name cannot be empty")
)

// Currency represents a currency with a code, a name, and the decimal places
// used to get the amount of a currency.
type Currency struct {
	code          CurrencyCode
	name          string
	decimalPlaces uint8
}

// NewCurrency expects a code of 3 capital letters, a name for the currency, and how many decimal places to
// use when calculating the amount, and it returns a Currency or an error.
func NewCurrency(code string, name string, decimalPlaces uint8) (Currency, error) {
	currencyCode := CurrencyCode(strings.ToUpper(strings.TrimSpace(code)))
	name = strings.TrimSpace(name)

	if !currencyCode.IsValid() {
		return Currency{}, ErrInvalidCurrencyCode
	}

	if len(name) < 1 {
		return Currency{}, ErrEmptyCurrencyName
	}

	return Currency{
		code:          currencyCode,
		name:          name,
		decimalPlaces: decimalPlaces,
	}, nil
}

// Code returns the 3 letter currency code.
func (c Currency) Code() CurrencyCode {
	return c.code
}

// Name returns the currency name.
func (c Currency) Name() string {
	return c.name
}

// DecimalPlaces returns the decimal places used to calculate the amount for a currency.
func (c Currency) DecimalPlaces() uint8 {
	return c.decimalPlaces
}

// String returns a string with this format "<CODE>: <Name>".
func (c Currency) String() string {
	return fmt.Sprintf("%s: %s", c.code, c.name)
}
