package currency_test

import (
	"path/filepath"
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/currency"
	currencyRepo "github.com/Mr-MichaelK/money-man/internal/infra/currency"
)

func TestNewSQLiteCurrencyRepository(t *testing.T) {
	const invalidDbPath = "/nonexistent-dir/test.db"

	t.Run("should return an error when given an invalid db path", func(t *testing.T) {
		_, err := currencyRepo.NewSQLiteCurrencyRepository(invalidDbPath)

		if err == nil {
			t.Fatal("expected an error when called Open in NewSQLiteCurrencyRepository but no error was returned")
		}
	})

	t.Run("should create the currencies table on initialization and return a repository", func(t *testing.T) {
		repo := initRepo(t)

		if repo == nil {
			t.Fatal("expected an NewSQLiteCurrencyRepository to be returned but nil was returned")
		}

		exists := currencyTableExistsHelper(t, repo)
		if !exists {
			t.Fatal("expected currencies table to be created but it was not")
		}
	})

	t.Run("should not fail when currencies table is created more than once", func(t *testing.T) {
		dbPath := filepath.Join(t.TempDir(), "test.db")

		repo, err := currencyRepo.NewSQLiteCurrencyRepository(dbPath)
		if err != nil {
			t.Fatalf("unexpected error when initializing repo for the first time: %v", err)
		}
		t.Cleanup(func() {
			if err := repo.DB().Close(); err != nil {
				t.Errorf("failed to close first repo db: %v", err)
			}
		})

		secondRepo, err := currencyRepo.NewSQLiteCurrencyRepository(dbPath)
		if err != nil {
			t.Fatalf("unexpected error when initializing repo for the second time: %v", err)
		}
		t.Cleanup(func() {
			if err := secondRepo.DB().Close(); err != nil {
				t.Errorf("failed to close second repo db: %v", err)
			}
		})

		exists := currencyTableExistsHelper(t, secondRepo)
		if !exists {
			t.Fatal("expected currencies table to exist after re-initializing repository")
		}
	})
}

func TestGetCurrencyByCurrencyCode(t *testing.T) {
	t.Run("should return an error if the currency does not exist", func(t *testing.T) {
		repo := initRepo(t)
		curr := getCurrency(t)

		_, err := repo.GetCurrencyByCurrencyCode(curr.Code())

		if err == nil {
			t.Fatal("expected an error to be returned but no error was returned")
		}
	})

	t.Run("should get a currency by its currency code and return it", func(t *testing.T) {
		repo := initRepo(t)
		savedCurrency := getCurrency(t)
		insertHelper(t, repo, &savedCurrency)

		curr, err := repo.GetCurrencyByCurrencyCode(savedCurrency.Code())

		if err != nil {
			t.Fatal("unexpected error was returned by GetCurrencyByCurrencyCode")
		}

		if curr != savedCurrency {
			t.Fatalf("expected %v but got %v", savedCurrency, curr)
		}

	})
}

func initRepo(t *testing.T) *currencyRepo.SQLiteCurrencyRepository {
	t.Helper()

	repo, err := currencyRepo.NewSQLiteCurrencyRepository(":memory:")
	if err != nil {
		t.Fatal("unexpected error was returned when initializing repo")
	}

	t.Cleanup(func() {
		if err := repo.DB().Close(); err != nil {
			t.Errorf("failed to close db: %v", err)
		}
	})

	return repo
}

func currencyTableExistsHelper(t *testing.T, repo *currencyRepo.SQLiteCurrencyRepository) bool {
	t.Helper()

	query := `SELECT COUNT(1) FROM sqlite_master WHERE type='table' AND name='currencies';`

	var count int
	err := repo.DB().QueryRow(query).Scan(&count)
	if err != nil {
		t.Fatal("unexpected error was returned when checking for existence of currencies table")
	}

	return count == 1
}

func insertHelper(t *testing.T, repo *currencyRepo.SQLiteCurrencyRepository, c *currency.Currency) {
	t.Helper()

	_, err := repo.DB().Exec(
		"INSERT INTO currencies (code, name, decimal_places) VALUES (?, ?, ?)",
		c.Code(),
		c.Name(),
		c.DecimalPlaces(),
	)

	if err != nil {
		t.Fatal("unexpected error was returned when inserting currency")
	}
}

func getCurrency(t *testing.T) currency.Currency {
	t.Helper()

	curr, err := currency.NewCurrency("USD", "dollar", 2)
	if err != nil {
		t.Fatalf("unexpected error was returned when creating currency %s", "USD")
	}

	return curr
}
