package currency

import (
	"errors"
	"fmt"
)

var ErrNilFetchCurrencyRepository = errors.New("fetch currency repository cannot be nil")

// FetchCurrencyRepository fetches a currency from its currency code.
type FetchCurrencyRepository interface {
	GetCurrencyByCurrencyCode(currencyCode CurrencyCode) (Currency, error)
}

// FetchCurrencyService takes a currency code and fetches the corresponding currency.
type FetchCurrencyService struct {
	repo FetchCurrencyRepository
}

// NewFetchCurrencyService takes a FetchCurrencyRepository and returns a FetchCurrencyService.
func NewFetchCurrencyService(fetchCurrencyRepository FetchCurrencyRepository) (FetchCurrencyService, error) {
	if fetchCurrencyRepository == nil {
		return FetchCurrencyService{}, ErrNilFetchCurrencyRepository
	}

	return FetchCurrencyService{repo: fetchCurrencyRepository}, nil
}

// Fetch takes a currency code and invokes the repository to fetch the currency.
func (s *FetchCurrencyService) Fetch(currencyCode CurrencyCode) (Currency, error) {
	curr, err := s.repo.GetCurrencyByCurrencyCode(currencyCode)
	if err != nil {
		return Currency{}, fmt.Errorf("fetch currency repository returned an error: %w", err)
	}

	return curr, nil
}
