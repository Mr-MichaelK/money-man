package currency

// CurrencyFetcher takes a CurrencyCode and returns its corresponding currency.
type CurrencyFetcher interface {
	Fetch(currencyCode CurrencyCode) (Currency, error)
}
