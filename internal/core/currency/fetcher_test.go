package currency_test

import (
	"errors"
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/currency"
)

type MockFetchCurrencyRepository struct {
	TimesGetCalled int

	MockCurrency      currency.Currency
	FetchReturnsError error
}

func (r *MockFetchCurrencyRepository) GetCurrencyByCurrencyCode(currencyCode currency.CurrencyCode) (currency.Currency, error) {
	r.TimesGetCalled++
	return r.MockCurrency, r.FetchReturnsError
}

func TestNewFetchCurrencyService(t *testing.T) {
	t.Run("should reject creating a FetchCurrencyService without a FetchCurrencyRepository", func(t *testing.T) {
		_, err := currency.NewFetchCurrencyService(nil)

		if err == nil {
			t.Fatal("expected an error to be returned but none were returned")
		}
	})

	t.Run("given a FetchCurrencyRepository the system should create a FetchCurrencyService", func(t *testing.T) {
		repo := &MockFetchCurrencyRepository{}
		_, err := currency.NewFetchCurrencyService(repo)

		if err != nil {
			t.Fatalf("unexpected error was returned when creating FetchCurrencyService")
		}
	})
}

func TestFetchCurrencyServiceFetch(t *testing.T) {
	const (
		invalidCurrencyCode = "2j9rp"
		currencyCode        = "USD"
		currencyName        = "dollar"
		decimalPlaces       = 2
	)

	t.Run("given an invalid currency code then the system should reject the request", func(t *testing.T) {
		repo := &MockFetchCurrencyRepository{}
		s, err := currency.NewFetchCurrencyService(repo)
		if err != nil {
			t.Fatal("unexpected error was returned when creating FetchCurrencyService")
		}

		_, err = s.Fetch(currency.CurrencyCode(invalidCurrencyCode))

		if err == nil {
			t.Fatal("expected an error to be returned but none were returned")
		}
	})

	t.Run("given the FetchCurrencyRepository returned an error then the system should propagate that error", func(t *testing.T) {
		repo := &MockFetchCurrencyRepository{FetchReturnsError: errors.New("error")}
		s, err := currency.NewFetchCurrencyService(repo)
		if err != nil {
			t.Fatal("unexpected error was returned when creating FetchCurrencyService")
		}

		_, err = s.Fetch(currency.CurrencyCode(currencyCode))

		if err == nil {
			t.Fatal("expected an error to be returned but none were returned")
		}
	})

	t.Run("should fetch and return the currency", func(t *testing.T) {
		mockCurrency, err := currency.NewCurrency(currencyCode, currencyName, decimalPlaces)
		if err != nil {
			t.Fatal("unexpected error was returned when creating currency")
		}

		repo := &MockFetchCurrencyRepository{MockCurrency: mockCurrency}

		s, err := currency.NewFetchCurrencyService(repo)
		if err != nil {
			t.Fatal("unexpected error was returned when creating FetchCurrencyService")
		}

		curr, err := s.Fetch(currency.CurrencyCode(currencyCode))

		if err != nil {
			t.Fatal("unexpected error was returned when Fetch was called")
		}

		if repo.TimesGetCalled != 1 {
			t.Fatalf("expected repo.GetCurrencyByCurrencyCode to be called once but was actually called %d times", repo.TimesGetCalled)
		}

		if curr != mockCurrency {
			t.Fatalf("expected %v but got %v", mockCurrency, curr)
		}
	})
}
