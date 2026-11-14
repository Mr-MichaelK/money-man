package profile_test

import (
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/profile"
)

func TestNew(t *testing.T) {
	const (
		username = "Michael"
		password = "supersecretpassword"
	)

	t.Run("should create a blank profile if no options are passed", func(t *testing.T) {
		want := profile.Profile{}
		got := *profile.New()

		if want != got {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})

	t.Run("should take a username option and initialize the profile with it", func(t *testing.T) {
		want := username
		got := profile.New(profile.WithUsername(username))

		if want != got.Username() {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})

	t.Run("should take a password option and initialize the profile with it", func(t *testing.T) {
		want := password
		got := profile.New(profile.WithPassword(password))

		if want != got.Password() {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})

	t.Run("should take multiple options and initialize the profile with it", func(t *testing.T) {
		want := profile.Profile{}
		want.SetUsername(username)
		want.SetPassword(password)

		got := *profile.New(
			profile.WithUsername(username),
			profile.WithPassword(password),
		)

		if want != got {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})
}
