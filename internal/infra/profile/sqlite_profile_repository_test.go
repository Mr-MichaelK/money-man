package profile_test

import (
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/profile"
	profileRepo "github.com/Mr-MichaelK/money-man/internal/infra/profile"
)

const (
	username = "Michael"
	password = "supersecretpassword"
)

func TestNewSQLiteProfileRepository(t *testing.T) {
	const invalidDbPath = "/nonexistent-dir/test.db"

	t.Run("should return an error when given an invalid db path", func(t *testing.T) {
		_, err := profileRepo.NewSQLiteProfileRepository(invalidDbPath)

		if err == nil {
			t.Fatal("expected an error when called Open in NewSQLiteRepository but no error was thrown")
		}
	})

	t.Run("should create the profiles table on initialization and return a repository", func(t *testing.T) {
		repo, err := initRepo()
		if err != nil {
			t.Fatal("unexpected error was returned when initializing repo")
		}

		if repo == nil {
			t.Fatal("expected an SQLiteRepository to be returned but nil was received")
		}

		exists, err := profileTableExistsHelper(repo)
		if err != nil {
			t.Fatal("unexpected error was returned when checking for existence of profiles table")
		}

		if !exists {
			t.Fatal("expected profiles table to be created but it was not")
		}
	})

	t.Run("should not fail when profiles table is created more than once", func(t *testing.T) {
		repo, err := initRepo()
		if err != nil {
			t.Fatal("unexpected error was returned when initializing repo for the first time")
		}

		_, err = initRepo()
		if err != nil {
			t.Fatal("unexpected error was returned when initializing repo for the second time")
		}

		exists, err := profileTableExistsHelper(repo)
		if err != nil {
			t.Fatal("unexpected error was returned when checking for existence of profiles table")
		}

		if !exists {
			t.Fatal("expected profiles table to be created but it was not")
		}
	})
}

func TestSave(t *testing.T) {
	t.Run("should insert a profile", func(t *testing.T) {
		repo, err := initRepo()
		if err != nil {
			t.Fatal("unexpected error was returned when initializing repo")
		}

		p := getProfile()

		err = repo.Save(p)
		if err != nil {
			t.Fatalf("unexpected error was returned when Save was called with %v", p)
		}

		exists, err := existsHelper(repo, username)
		if err != nil {
			t.Fatal("unexpected error was returned when checking for existence of profile")
		}

		if !exists {
			t.Errorf("expected profile %v to be inserted but it does not exist", p)
		}
	})

	t.Run("should reject the request if the profile already exists", func(t *testing.T) {
		repo, err := initRepo()
		if err != nil {
			t.Fatal("unexpected error was returned when initializing repo")
		}

		p := getProfile()

		err = insertHelper(repo, p)
		if err != nil {
			t.Fatal("unexpected error was returned when inserting profile")
		}

		err = repo.Save(p)
		if err == nil {
			t.Fatalf("expected an error to be returned when Save was called with %v a second time", p)
		}
	})
}

func TestExistsByUsername(t *testing.T) {
	t.Run("should return false when profile does not exist", func(t *testing.T) {
		repo, err := initRepo()
		if err != nil {
			t.Fatal("unexpected error was returned when initializing repo")
		}

		exists, err := repo.ExistsByUsername(username)
		if err != nil {
			t.Fatal("unexpected error was returned when checking for existence of profile")
		}

		if exists {
			t.Error("expected profile to not exist but it did")
		}
	})

	t.Run("should return true when profile exists", func(t *testing.T) {
		repo, err := initRepo()
		if err != nil {
			t.Fatal("unexpected error was returned when initializing repo")
		}

		p := getProfile()

		err = insertHelper(repo, p)
		if err != nil {
			t.Fatal("unexpected error was returned when inserting profile")
		}

		exists, err := repo.ExistsByUsername(username)
		if err != nil {
			t.Fatal("unexpected error was returned when checking for existence of profile")
		}

		if !exists {
			t.Error("expected profile to exist but it did not")
		}
	})
}

func initRepo() (*profileRepo.SQLiteProfileRepository, error) {
	return profileRepo.NewSQLiteProfileRepository(":memory:")
}

func insertHelper(repo *profileRepo.SQLiteProfileRepository, p *profile.Profile) error {
	_, err := repo.DB().Exec(
		"INSERT INTO profiles (id, username, password) VALUES (?, ?, ?)",
		p.Id(),
		p.Username(),
		p.Password(),
	)

	return err
}

func existsHelper(repo *profileRepo.SQLiteProfileRepository, username string) (bool, error) {
	query := `SELECT COUNT(1) FROM profiles WHERE username = ?`
	var count int

	err := repo.DB().QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func profileTableExistsHelper(repo *profileRepo.SQLiteProfileRepository) (bool, error) {
	query := `SELECT COUNT(1) FROM sqlite_master WHERE type='table' AND name='profiles';`

	var count int
	err := repo.DB().QueryRow(query).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func getProfile() *profile.Profile {
	return profile.New(
		profile.WithUsername(username),
		profile.WithPassword(password),
	)
}
