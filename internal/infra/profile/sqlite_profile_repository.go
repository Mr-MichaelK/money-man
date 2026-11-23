package profile

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mr-MichaelK/money-man/internal/core/profile"

	// Register the pure-Go SQLite driver with database/sql.
	_ "modernc.org/sqlite"
)

var (
	ErrFailedToOpenSQLite               = errors.New("failed to open sqlite db")
	ErrFailedToPingSQLite               = errors.New("failed to ping sqlite db")
	ErrFailedToCreateTable              = errors.New("failed to create profiles table")
	ErrFailedToSaveProfile              = errors.New("failed to save profile")
	ErrFailedToCheckExistenceByUsername = errors.New("failed to check username existence")
)

const (
	createTableQuery = `
	CREATE TABLE IF NOT EXISTS profiles (
		id TEXT PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	insertProfileQuery    = `INSERT INTO profiles (id, username, password) VALUES (?, ?, ?)`
	existsByUsernameQuery = `SELECT COUNT(1) FROM profiles WHERE username = ?`
)

// SQLiteProfileRepository defines the methods to interact with the profiles table in SQLite.
type SQLiteProfileRepository struct {
	db *sql.DB
}

// DB is a getter which returns SQLiteProfileRepository.db field.
func (r *SQLiteProfileRepository) DB() *sql.DB {
	return r.db
}

// NewSQLiteProfileRepository initializes the repository and ensures the profiles table exists.
func NewSQLiteProfileRepository(dbPath string) (*SQLiteProfileRepository, error) {
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

	repo := &SQLiteProfileRepository{db: db}
	if err := repo.createTable(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToCreateTable, err)
	}

	success = true
	return repo, nil
}

// createTable creates the "profiles" table.
func (r *SQLiteProfileRepository) createTable() error {
	query := createTableQuery
	_, err := r.db.Exec(query)
	return err
}

// Save inserts a profile into SQLite.
func (r *SQLiteProfileRepository) Save(p *profile.Profile) error {
	query := insertProfileQuery

	_, err := r.db.Exec(query, p.Id(), p.Username(), p.Password())
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedToSaveProfile, err)
	}
	return nil
}

// ExistsByUsername checks if a username is already taken.
func (r *SQLiteProfileRepository) ExistsByUsername(username string) (bool, error) {
	query := existsByUsernameQuery

	var count int
	err := r.db.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("%w: %w", ErrFailedToCheckExistenceByUsername, err)
	}

	return count > 0, nil
}
