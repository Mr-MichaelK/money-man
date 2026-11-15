package profile_test

import (
	"testing"
	"uuid"

	"github.com/Mr-MichaelK/money-man/internal/core/profile"
)

func TestNew(t *testing.T) {
	const (
		username = "Michael"
		password = "supersecretpassword"
	)

	t.Run("should create a blank profile with an id if no options are passed", func(t *testing.T) {
		want := profile.Profile{}

		got := profile.New()

		if _, err := uuid.Parse(got.Id()); err != nil {
			t.Errorf("expected ID to be a valid UUID, got %v", got.Id())
		}

		if !isEqualProfile(want, *got) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})

	t.Run("should take a username option and initialize the profile with it", func(t *testing.T) {
		want := profile.Profile{}
		want.SetUsername(username)

		got := profile.New(profile.WithUsername(username))

		if !isEqualProfile(want, *got) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})

	t.Run("should take a password option and initialize the profile with it", func(t *testing.T) {
		want := profile.Profile{}
		want.SetPassword(password)

		got := profile.New(profile.WithPassword(password))

		if !isEqualProfile(want, *got) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})

	t.Run("should take multiple options and initialize the profile with it", func(t *testing.T) {
		want := profile.Profile{}
		want.SetUsername(username)
		want.SetPassword(password)

		got := profile.New(
			profile.WithUsername(username),
			profile.WithPassword(password),
		)

		if !isEqualProfile(want, *got) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})
}

func isEqualProfile(p1, p2 profile.Profile) bool {
	return p1.Username() == p2.Username() &&
		p1.Password() == p2.Password()
}
