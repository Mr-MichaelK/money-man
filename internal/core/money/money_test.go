package money_test

import (
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/money"
)

func TestMoneyString(t *testing.T) {
	t.Run("should return the amount of money with its decimal places and the currency code beside it", func(t *testing.T) {
		amount := money.Amount(10000)
		currency, err := money.NewCurrency("USD", "dollar", 2)
		if err != nil {
			t.Fatal("unexpected error occured when creating currency")
		}

		m := money.New(amount, currency)
		got := m.String()

		expected := "100.00 USD"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})

	t.Run("should format negative amounts smaller than one unit with leading zero", func(t *testing.T) {
		amount := money.Amount(-5)
		currency, err := money.NewCurrency("USD", "US Dollar", 2)
		if err != nil {
			t.Fatalf("unexpected error occurred when creating currency: %v", err)
		}

		m := money.New(amount, currency)
		got := m.String()

		expected := "-0.05 USD"
		if got != expected {
			t.Errorf("expected %s but got %s", expected, got)
		}
	})
}

func TestMoneyAdd(t *testing.T) {
	var (
		amountOne = money.Amount(30)
		amountTwo = money.Amount(20)

		currencyCodeOne = "AAA"
		currencyNameOne = "a"

		currencyCodeTwo = "BBB"
		currencyNameTwo = "b"
	)

	getCurrencyOne := func() money.Currency {
		c, err := money.NewCurrency(currencyCodeOne, currencyNameOne, 0)
		if err != nil {
			t.Fatal("unexpected error occured when creating currency")
		}

		return c
	}

	getCurrencyTwo := func() money.Currency {
		c, err := money.NewCurrency(currencyCodeTwo, currencyNameTwo, 0)
		if err != nil {
			t.Fatal("unexpected error occured when creating currency")
		}

		return c
	}

	getMoneyOne := func() money.Money {
		return money.New(amountOne, getCurrencyOne())
	}

	getMoneyTwo := func() money.Money {
		return money.New(amountTwo, getCurrencyTwo())
	}

	t.Run("should reject adding two amounts of money with different currencies", func(t *testing.T) {
		m1 := getMoneyOne()
		m2 := getMoneyTwo()

		_, err := m1.Add(m2)

		if err == nil {
			t.Fatalf("expected an error to be returned when adding currencies \"%s\" and \"%s\" but no error was returned", currencyCodeOne, currencyCodeTwo)
		}
	})

	t.Run("should add two amounts of money with the same currency", func(t *testing.T) {
		m1 := getMoneyOne()
		m2 := getMoneyOne()

		result, err := m1.Add(m2)

		if err != nil {
			t.Fatalf("unexpected error occured when adding %v and %v", m1, m2)
		}

		expectedAmount := money.Amount(60)
		if result.Amount() != expectedAmount {
			t.Fatalf("expected %v but got %v", expectedAmount, result.Amount())
		}
	})
}

func TestMoneySubtract(t *testing.T) {
	var (
		amountOne = money.Amount(30)
		amountTwo = money.Amount(20)

		currencyCodeOne = "AAA"
		currencyNameOne = "a"

		currencyCodeTwo = "BBB"
		currencyNameTwo = "b"
	)

	getCurrencyOne := func() money.Currency {
		c, err := money.NewCurrency(currencyCodeOne, currencyNameOne, 0)
		if err != nil {
			t.Fatal("unexpected error occured when creating currency")
		}

		return c
	}

	getCurrencyTwo := func() money.Currency {
		c, err := money.NewCurrency(currencyCodeTwo, currencyNameTwo, 0)
		if err != nil {
			t.Fatal("unexpected error occured when creating currency")
		}

		return c
	}

	getMoneyOne := func() money.Money {
		return money.New(amountOne, getCurrencyOne())
	}

	getMoneyTwo := func() money.Money {
		return money.New(amountTwo, getCurrencyTwo())
	}

	t.Run("should reject subtracting two amounts of money with different currencies", func(t *testing.T) {
		m1 := getMoneyOne()
		m2 := getMoneyTwo()

		_, err := m1.Subtract(m2)

		if err == nil {
			t.Fatalf("expected an error to be returned when subtracting currencies \"%s\" and \"%s\" but no error was returned", currencyCodeOne, currencyCodeTwo)
		}
	})

	t.Run("should subtract two amounts of money with the same currency", func(t *testing.T) {
		m1 := getMoneyOne()
		m2 := getMoneyOne()

		result, err := m1.Subtract(m2)

		if err != nil {
			t.Fatalf("unexpected error occured when subtracting %v and %v", m1, m2)
		}

		expectedAmount := money.Amount(0)
		if result.Amount() != expectedAmount {
			t.Fatalf("expected %v but got %v", expectedAmount, result.Amount())
		}
	})
}
