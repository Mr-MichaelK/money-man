package profile_test

import (
	"errors"
	"testing"

	"github.com/Mr-MichaelK/money-man/internal/core/profile"
)

const (
	username = "Michael"
	password = "supersecretpassword"
	id       = "1234"
)

type MockRegistrationRepository struct {
	SaveCalledCount   int
	ExistsCalledCount int

	ExistsReturnBool  bool
	ExistsReturnError error
	SaveReturnError   error
}

func (repo *MockRegistrationRepository) Save(p *profile.Profile) error {
	repo.SaveCalledCount++
	return repo.SaveReturnError
}

func (repo *MockRegistrationRepository) ExistsByUsername(username string) (bool, error) {
	repo.ExistsCalledCount++
	return repo.ExistsReturnBool, repo.ExistsReturnError
}

func TestRegister(t *testing.T) {
	t.Run("should check if the username exists before saving", func(t *testing.T) {
		repo := &MockRegistrationRepository{}
		registrationService := profile.NewRegistrationService(repo)

		registrationService.Register(getRegisterProfileRequest())

		expectedTimesCalled := 1
		if repo.ExistsCalledCount != expectedTimesCalled {
			t.Errorf("expected \"ExistsByUsername\" to be called %d times but instead it was called %d times", expectedTimesCalled, repo.ExistsCalledCount)
		}
	})

	t.Run("should reject the request if an error occurs when checking for username existence", func(t *testing.T) {
		repo := &MockRegistrationRepository{
			ExistsReturnError: errors.New("database connection failed"),
		}
		registrationService := profile.NewRegistrationService(repo)

		_, err := registrationService.Register(getRegisterProfileRequest())

		if err == nil {
			t.Error("expected \"ExistsByUsername\" to return an error but no error was thrown")
		}

		if repo.SaveCalledCount != 0 {
			t.Errorf("expected \"ExistsByUsername\" to not invoke repo.Save but it got invoked %d times", repo.SaveCalledCount)
		}
	})

	t.Run("should reject the request if the username already exists", func(t *testing.T) {
		repo := &MockRegistrationRepository{
			ExistsReturnBool: true,
		}
		registrationService := profile.NewRegistrationService(repo)

		_, err := registrationService.Register(getRegisterProfileRequest())

		if err == nil {
			t.Error("expected \"ExistsByUsername\" to return an error but no error was thrown")
		}

		if repo.SaveCalledCount != 0 {
			t.Errorf("expected \"ExistsByUsername\" to not invoke repo.Save but it got invoked %d times", repo.SaveCalledCount)
		}
	})

	t.Run("should create and save a profile", func(t *testing.T) {
		repo := &MockRegistrationRepository{}
		registrationService := profile.NewRegistrationService(repo)

		registrationService.Register(getRegisterProfileRequest())

		expectedTimesCalled := 1
		if repo.SaveCalledCount != expectedTimesCalled {
			t.Errorf("expected \"Save\" to be called %d times but instead it was called %d times", expectedTimesCalled, repo.SaveCalledCount)
		}
	})

	t.Run("should reject the request if an error occurs when saving", func(t *testing.T) {
		repo := &MockRegistrationRepository{
			SaveReturnError: errors.New("database connection failed"),
		}
		registrationService := profile.NewRegistrationService(repo)

		_, err := registrationService.Register(getRegisterProfileRequest())

		if err == nil {
			t.Error("expected \"Save\" to return an error but no error was thrown")
		}
	})

	t.Run("should create a profile, save it, and return a response", func(t *testing.T) {
		repo := &MockRegistrationRepository{}
		registrationService := profile.NewRegistrationService(repo)

		response, err := registrationService.Register(getRegisterProfileRequest())

		if err != nil {
			t.Fatal("was not expecting an error but an error was returned")
		}

		expectedResponse := getExpectedRegisterProfileResponse()
		if !isEqualRegisterProfileResponse(response, expectedResponse) {
			t.Fatalf("wanted %v got %v", expectedResponse, response)
		}
	})
}

func getRegisterProfileRequest() profile.RegisterProfileRequest {
	return profile.RegisterProfileRequest{
		Username: username,
		Password: password,
	}
}

func getExpectedRegisterProfileResponse() profile.RegisterProfileResponse {
	return profile.RegisterProfileResponse{
		Id:       id,
		Username: username,
	}
}

func isEqualRegisterProfileResponse(r1, r2 profile.RegisterProfileResponse) bool {
	return r1.Username == r2.Username
}
