package money

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrInvalidCurrencyCode = errors.New("currency code should be 3 capital letters")
	ErrEmptyCurrencyName   = errors.New("currency name cannot be empty")
)

// CurrencyCode is a string of length 3 such as USD, EUR, and LBP.
type CurrencyCode string

// Currency represents a currency with a code, a name, and the decimal places
// used to get the amount of a currency.
type Currency struct {
	code          CurrencyCode
	name          string
	decimalPlaces uint8
}

var currencyCodeRegex = regexp.MustCompile(`^[A-Z]{3}$`)

// NewCurrency expects a code of 3 capital letters, a name for the currency, and how many decimal places to
// use when calculating the amount, and it returns a Currency or an error.
func NewCurrency(code string, name string, decimalPlaces uint8) (Currency, error) {
	code = strings.ToUpper(strings.TrimSpace(code))
	name = strings.TrimSpace(name)

	if !currencyCodeRegex.MatchString(code) {
		return Currency{}, ErrInvalidCurrencyCode
	}

	if len(name) < 1 {
		return Currency{}, ErrEmptyCurrencyName
	}

	return Currency{
		code:          CurrencyCode(code),
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

// String returns a string with this format "<CODE>: <Name>".
func (c Currency) String() string {
	return fmt.Sprintf("%s: %s", c.code, c.name)
}
