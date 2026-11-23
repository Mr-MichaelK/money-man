package currency

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mr-MichaelK/money-man/internal/core/currency"

	// Register the pure-Go SQLite driver with database/sql.
	_ "modernc.org/sqlite"
)

var (
	ErrFailedToOpenSQLite  = errors.New("failed to open sqlite db")
	ErrFailedToPingSQLite  = errors.New("failed to ping sqlite db")
	ErrFailedToCreateTable = errors.New("failed to create currencies table")
	ErrFailedToGetCurrency = errors.New("failed to get currency")
	ErrCurrencyNotFound    = errors.New("currency not found")
)

const (
	createTableQuery = `
	CREATE TABLE IF NOT EXISTS currencies (
		code TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		decimal_places INT DEFAULT 0
	);`
	getCurrencyByCurrencyCodeQuery = `
	SELECT code, name, decimal_places
	FROM currencies
	WHERE code = ?`
)

// SQLiteCurrencyRepository defines the methods to interact with the currencies table in SQLite.
type SQLiteCurrencyRepository struct {
	db *sql.DB
}

// DB is a getter which returns SQLiteCurrencyRepository.db field.
func (r *SQLiteCurrencyRepository) DB() *sql.DB {
	return r.db
}

// NewSQLiteCurrencyRepository initializes the repository and ensures the currencies table exists.
func NewSQLiteCurrencyRepository(dbPath string) (*SQLiteCurrencyRepository, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToOpenSQLite, err)
	}

	success := false
	defer func() {
		if !success {
			db.Close()
		}
	}()

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToPingSQLite, err)
	}

	repo := &SQLiteCurrencyRepository{db: db}
	if err := repo.createTable(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToCreateTable, err)
	}

	success = true
	return repo, nil
}

// createTable creates the "currencies" table.
func (r *SQLiteCurrencyRepository) createTable() error {
	query := createTableQuery
	_, err := r.db.Exec(query)
	return err
}

// GetCurrencyByCurrencyCode finds the Currency by CurrencyCode in SQLite.
func (r *SQLiteCurrencyRepository) GetCurrencyByCurrencyCode(currencyCode currency.CurrencyCode) (currency.Currency, error) {
	var (
		code          string
		name          string
		decimalPlaces uint8
	)

	err := r.db.QueryRow(getCurrencyByCurrencyCodeQuery, currencyCode).Scan(
		&code,
		&name,
		&decimalPlaces,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return currency.Currency{}, fmt.Errorf("%w %s: %w", ErrCurrencyNotFound, currencyCode, err)
		}
		return currency.Currency{}, fmt.Errorf("%w %s: %w", ErrFailedToGetCurrency, currencyCode, err)
	}

	return currency.NewCurrency(code, name, decimalPlaces)
}
