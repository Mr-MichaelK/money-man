package money_test

import (
	"errors"
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/currency"
	"github.com/Mr-MichaelK/money-man/internal/core/money"
)

type MockCurrencyFetcher struct {
	MockCurrency      currency.Currency
	FetchReturnsError error
}

func (mock MockCurrencyFetcher) Fetch(currencyCode currency.CurrencyCode) (currency.Currency, error) {
	return mock.MockCurrency, mock.FetchReturnsError
}

func TestNewMoneyFactory(t *testing.T) {
	t.Run("should reject creating a money factory without a currency fetcher", func(t *testing.T) {
		_, err := money.NewMoneyFactory(nil)
		if err == nil {
			t.Fatal("expected an error but no error was returned")
		}
	})

	t.Run("should create a money factory with a currency fetcher", func(t *testing.T) {
		currencyFetcher := MockCurrencyFetcher{FetchReturnsError: errors.New("error")}
		_, err := money.NewMoneyFactory(currencyFetcher)
		if err != nil {
			t.Fatal("unexpected error when creating money factory")
		}
	})
}

func TestMoneyFactoryNewMoney(t *testing.T) {
	t.Run("should reject to create money if the currency code is invalid", func(t *testing.T) {
		currencyFetcher := MockCurrencyFetcher{}
		moneyFactory, err := money.NewMoneyFactory(currencyFetcher)
		if err != nil {
			t.Fatal("unexpected error when creating money factory")
		}

		_, err = moneyFactory.NewMoney(money.Amount(10), currency.CurrencyCode("1!a"))
		if err == nil {
			t.Fatal("expected an error but no error was returned")
		}
	})

	t.Run("given the system cannot fetch the currency then the system should reject creating money", func(t *testing.T) {
		currencyFetcher := MockCurrencyFetcher{FetchReturnsError: errors.New("error")}
		moneyFactory, err := money.NewMoneyFactory(currencyFetcher)
		if err != nil {
			t.Fatal("unexpected error when creating money factory")
		}

		_, err = moneyFactory.NewMoney(money.Amount(10), currency.CurrencyCode("USD"))

		if err == nil {
			t.Fatal("expected an error to be returned but no error was returned")
		}
	})

	t.Run("should create money", func(t *testing.T) {
		mockCurrency, err := currency.NewCurrency("USD", "dollar", 2)
		if err != nil {
			t.Fatal("unexpected error was returned when creating currency")
		}

		currencyFetcher := MockCurrencyFetcher{MockCurrency: mockCurrency}
		moneyFactory, err := money.NewMoneyFactory(currencyFetcher)
		if err != nil {
			t.Fatal("unexpected error when creating money factory")
		}

		m, err := moneyFactory.NewMoney(money.Amount(10), currency.CurrencyCode("USD"))
		if err != nil {
			t.Fatal("unexpected error was returned when creating money")
		}

		expectedMoney := money.New(money.Amount(10), mockCurrency)
		if m != expectedMoney {
			t.Fatalf("expected %v but got %v", expectedMoney, m)
		}
	})
}
