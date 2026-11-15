package profile

import "github.com/Mr-MichaelK/money-man/internal/core/profile"

// TODO: Implement in a future commit
type SQLiteRepository struct{}

func Save(p *profile.Profile) error {
	return nil
}

func FindByUsername(username string) (profile.Profile, error) {
	return profile.Profile{}, nil
}
