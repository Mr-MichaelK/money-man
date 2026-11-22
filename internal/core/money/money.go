package money

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrCurrencyMismatch = errors.New("Currency mismatch: cannot operate on two amounts of money with different currencies")

// Money represents a certain amount of money with a currency.
type Money struct {
	amount   Amount
	currency Currency
}

// New takes an amount and currency and returns Money.
func New(amount Amount, currency Currency) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

// Amount is a m.amount getter.
func (m Money) Amount() Amount {
	return m.amount
}

// Currency is a m.currency getter.
func (m Money) Currency() Currency {
	return m.currency
}

// String returns the amount of money followed by the currency code.
func (m Money) String() string {
	raw := strconv.FormatInt(int64(m.amount), 10)

	places := int(m.currency.decimalPlaces)

	if places == 0 {
		return fmt.Sprintf("%s %s", raw, m.currency.Code())
	}

	sign := ""

	if strings.HasPrefix(raw, "-") {
		sign = "-"
		raw = raw[1:]
	}

	for len(raw) <= places {
		raw = "0" + raw
	}

	split := len(raw) - places

	return fmt.Sprintf("%s%s.%s %s", sign, raw[:split], raw[split:], m.currency.Code())
}

// Add takes a certain amount of money and adds it only if the currencies match and returns an error if not.
func (m Money) Add(other Money) (Money, error) {
	err := validateTransaction(m, other)
	if err != nil {
		return Money{}, err
	}

	m.amount.Add(other.amount)
	return m, nil
}

// Subtract takes a certain amount of money and subtraxts it only if the currencies match and returns an
// error if they do not.
func (m Money) Subtract(other Money) (Money, error) {
	err := validateTransaction(m, other)
	if err != nil {
		return Money{}, err
	}

	m.amount.Subtract(other.amount)
	return m, nil
}

// validateTransaction returns an error if the currencies of 2 Money instances do not match.
func validateTransaction(m1, m2 Money) error {
	if m1.Currency().Code() != m2.Currency().Code() {
		return ErrCurrencyMismatch
	}

	return nil
}
