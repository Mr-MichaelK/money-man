package currency

import "regexp"

var currencyCodeRegex = regexp.MustCompile(`^[A-Z]{3}$`)

// CurrencyCode is a string of length 3 such as USD, EUR, and LBP.
type CurrencyCode string

// IsValid determines whether a CurrencyCode is valid.
func (cc CurrencyCode) IsValid() bool {
	return currencyCodeRegex.MatchString(string(cc))
}
