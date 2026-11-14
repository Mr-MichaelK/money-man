package profile

import "uuid"

// Profile defines a single user profile with a username and password.
type Profile struct {
	id       string
	username string
	password string
}

// Option defines custom options for creating a new profile.
type Option func(*Profile)

// WithUsername creates an Option for creating a profile with a username.
func WithUsername(username string) Option {
	return func(p *Profile) {
		p.SetUsername(username)
	}
}

// WithPassword creates an Option for creating a profile with a password.
func WithPassword(password string) Option {
	return func(p *Profile) {
		p.SetPassword(password)
	}
}

// New takes options for creating a Profile and returns a pointer to a Profile.
func New(options ...Option) *Profile {
	profile := &Profile{id: uuid.New().String()}

	for _, option := range options {
		option(profile)
	}

	return profile
}

func (p *Profile) Id() string {
	return p.id
}

// Username is a Profile.username getter.
func (p *Profile) Username() string {
	return p.username
}

// Password is a Profile.profile getter.
func (p *Profile) Password() string {
	return p.password
}

// SetUsername takes a username string and assigns it to a profile.
func (p *Profile) SetUsername(username string) {
	p.username = username
}

// SetPassword takes a password string and assigns it to a profile.
func (p *Profile) SetPassword(password string) {
	p.password = password
}
