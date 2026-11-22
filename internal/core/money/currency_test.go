package money_test

import (
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/money"
)

func TestNewCurrency(t *testing.T) {
	const (
		emptyStr                          = ""
		invalidCodeWithSpaces             = " a "
		validCode                         = "AAA"
		invalidCode                       = "2!A"
		validCodeWithSpaces               = " " + validCode + "  "
		validCurrencyName                 = "name"
		validCurrencyNameWithSpaces       = "  " + validCurrencyName
		decimalPlaces               uint8 = 3
	)

	t.Run("should reject creating a currency with an empty code", func(t *testing.T) {
		_, err := money.NewCurrency(emptyStr, validCurrencyName, decimalPlaces)

		if err == nil {
			t.Fatalf("expected an error to be returned for currency code \"%s\" but no error was returned", emptyStr)
		}
	})

	t.Run("should reject creating a currency with code length less than 3 with spaces", func(t *testing.T) {
		_, err := money.NewCurrency(invalidCodeWithSpaces, validCurrencyName, decimalPlaces)

		if err == nil {
			t.Fatalf("expected an error to be returned for currency code \"%s\" but no error was returned", invalidCodeWithSpaces)
		}
	})

	t.Run("should reject creating a currency with a non-iso code", func(t *testing.T) {
		_, err := money.NewCurrency(invalidCode, validCurrencyName, decimalPlaces)

		if err == nil {
			t.Fatalf("expected an error to be returned for currency code \"%s\" but no error was returned", invalidCode)
		}
	})

	t.Run("should reject creating a currency with an empty name", func(t *testing.T) {
		_, err := money.NewCurrency(validCode, emptyStr, decimalPlaces)

		if err == nil {
			t.Fatalf("expected an error to be returned for currency name \"%s\" but no error was returned", emptyStr)
		}
	})

	t.Run("should reject creating a currency with blank name", func(t *testing.T) {
		_, err := money.NewCurrency(validCode, "  ", decimalPlaces)

		if err == nil {
			t.Fatalf("expected an error to be returned for currency name \"%s\" but no error was returned", "  ")
		}
	})

	t.Run("given the code and currency are valid but contain spaces, the system should trim the spaces and create the currency", func(t *testing.T) {
		c, err := money.NewCurrency(validCodeWithSpaces, validCurrencyNameWithSpaces, decimalPlaces)

		if err != nil {
			t.Fatalf("unexpected error was returned with currency code \"%s\" and currency name \"%s\"", validCodeWithSpaces, validCurrencyNameWithSpaces)
		}

		expectedCode := money.CurrencyCode(validCode)
		if c.Code() != expectedCode {
			t.Fatalf("expected currency code to be \"%s\" but got \"%s\"", expectedCode, c.Code())
		}

		expectedName := validCurrencyName
		if c.Name() != expectedName {
			t.Fatalf("expected currency name to be \"%s\" but got \"%s\"", expectedName, c.Name())
		}
	})

	t.Run("given the code and currency are valid then the system should create a currency using them", func(t *testing.T) {
		c, err := money.NewCurrency(validCode, validCurrencyName, decimalPlaces)

		if err != nil {
			t.Fatalf("unexpected error was returned with currency code \"%s\" and currency name \"%s\"", validCode, validCurrencyName)
		}

		expectedCode := money.CurrencyCode(validCode)
		if c.Code() != expectedCode {
			t.Fatalf("expected currency code to be \"%s\" but got \"%s\"", expectedCode, c.Code())
		}

		expectedName := validCurrencyName
		if c.Name() != expectedName {
			t.Fatalf("expected currency name to be \"%s\" but got \"%s\"", expectedName, c.Name())
		}
	})
}

func TestCurrencyString(t *testing.T) {
	const (
		code = "USD"
		name = "Dollars"
	)

	t.Run("should return a string formatted to have the code first followed by the name", func(t *testing.T) {
		c, err := money.NewCurrency(code, name, 2)
		if err != nil {
			t.Fatalf("unexpected error was returned with currency code \"%s\" and currency name \"%s\"", code, name)
		}

		formattedCurrency := c.String()

		expectedCurrency := code + ": " + name
		if formattedCurrency != expectedCurrency {
			t.Fatalf("expected %s but got %s", expectedCurrency, formattedCurrency)
		}
	})
}
