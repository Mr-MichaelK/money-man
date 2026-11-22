package currency

// CurrencyRegisterer registers a new Currency.
type CurrencyRegisterer interface {
	Register(currency Currency) error
}
