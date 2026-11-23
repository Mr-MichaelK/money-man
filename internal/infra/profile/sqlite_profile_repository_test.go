package profile_test

import (
	"path/filepath"
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
			t.Fatal("expected an error when called Open in NewSQLiteProfileRepository but no error was returned")
		}
	})

	t.Run("should create the profiles table on initialization and return a repository", func(t *testing.T) {
		repo := initRepo(t)

		if repo == nil {
			t.Fatal("expected an SQLiteProfileRepository to be returned but nil was returned")
		}

		exists := profileTableExistsHelper(t, repo)

		if !exists {
			t.Fatal("expected profiles table to be created but it was not")
		}
	})

	t.Run("should not fail when profiles table is created more than once", func(t *testing.T) {
		dbPath := filepath.Join(t.TempDir(), "test.db")

		repo, err := profileRepo.NewSQLiteProfileRepository(dbPath)
		if err != nil {
			t.Fatalf("unexpected error when initializing repo for the first time: %v", err)
		}
		t.Cleanup(func() {
			if err := repo.DB().Close(); err != nil {
				t.Errorf("failed to close first repo db: %v", err)
			}
		})

		secondRepo, err := profileRepo.NewSQLiteProfileRepository(dbPath)
		if err != nil {
			t.Fatalf("unexpected error when initializing repo for the second time: %v", err)
		}
		t.Cleanup(func() {
			if err := secondRepo.DB().Close(); err != nil {
				t.Errorf("failed to close second repo db: %v", err)
			}
		})

		exists := profileTableExistsHelper(t, secondRepo)
		if !exists {
			t.Fatal("expected profiles table to exist after re-initializing repository")
		}
	})
}

func TestSave(t *testing.T) {
	t.Run("should insert a profile", func(t *testing.T) {
		repo := initRepo(t)
		p := getProfile()

		err := repo.Save(p)
		if err != nil {
			t.Fatalf("unexpected error was returned when Save was called with %v", p)
		}

		exists := existsHelper(t, repo, username)

		if !exists {
			t.Errorf("expected profile %v to be inserted but it does not exist", p)
		}
	})

	t.Run("should reject the request if the profile already exists", func(t *testing.T) {
		repo := initRepo(t)
		p := getProfile()
		insertHelper(t, repo, p)

		err := repo.Save(p)
		if err == nil {
			t.Fatalf("expected an error to be returned when Save was called with %v a second time", p)
		}
	})
}

func TestExistsByUsername(t *testing.T) {
	t.Run("should return false when profile does not exist", func(t *testing.T) {
		repo := initRepo(t)

		exists, err := repo.ExistsByUsername(username)
		if err != nil {
			t.Fatal("unexpected error was returned when checking for existence of profile")
		}

		if exists {
			t.Error("expected profile to not exist but it did")
		}
	})

	t.Run("should return true when profile exists", func(t *testing.T) {
		repo := initRepo(t)
		p := getProfile()
		insertHelper(t, repo, p)

		exists, err := repo.ExistsByUsername(username)
		if err != nil {
			t.Fatal("unexpected error was returned when checking for existence of profile")
		}

		if !exists {
			t.Error("expected profile to exist but it did not")
		}
	})
}

func initRepo(t *testing.T) *profileRepo.SQLiteProfileRepository {
	t.Helper()

	repo, err := profileRepo.NewSQLiteProfileRepository(":memory:")
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

func insertHelper(t *testing.T, repo *profileRepo.SQLiteProfileRepository, p *profile.Profile) {
	t.Helper()

	_, err := repo.DB().Exec(
		"INSERT INTO profiles (id, username, password) VALUES (?, ?, ?)",
		p.Id(),
		p.Username(),
		p.Password(),
	)

	if err != nil {
		t.Fatal("unexpected error was returned when inserting profile")
	}
}

func existsHelper(t *testing.T, repo *profileRepo.SQLiteProfileRepository, username string) bool {
	t.Helper()

	query := `SELECT COUNT(1) FROM profiles WHERE username = ?`

	var count int

	err := repo.DB().QueryRow(query, username).Scan(&count)
	if err != nil {
		t.Fatal("unexpected error was returned when checking for existence of profile")
	}

	return count > 0
}

func profileTableExistsHelper(t *testing.T, repo *profileRepo.SQLiteProfileRepository) bool {
	t.Helper()

	query := `SELECT COUNT(1) FROM sqlite_master WHERE type='table' AND name='profiles';`

	var count int
	err := repo.DB().QueryRow(query).Scan(&count)
	if err != nil {
		t.Fatal("unexpected error was returned when checking for existence of profiles table")
	}

	return count == 1
}

func getProfile() *profile.Profile {
	return profile.New(
		profile.WithUsername(username),
		profile.WithPassword(password),
	)
}
