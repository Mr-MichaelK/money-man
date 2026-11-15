package profile

import "fmt"

const ErrUsernameExists = "Failed to register profile with username %s because it already exists"

// RegisterProfileRequest contains the information necessary to register a profile.
type RegisterProfileRequest struct {
	Username string
	Password string
}

// RegisterProfileResponse contains information about the profile.
type RegisterProfileResponse struct {
	Id       string
	Username string
}

// RegistrationRepository defines methods for interacting with the profile persistance layer.
type RegistrationRepository interface {
	Save(p *Profile) error
	ExistsByUsername(username string) (bool, error)
}

// RegistrationService coordinates the profile creation workflow.
type RegistrationService struct {
	repo RegistrationRepository
}

// NewRegistrationService initializes the service with its required repository dependency.
func NewRegistrationService(repo RegistrationRepository) *RegistrationService {
	return &RegistrationService{repo: repo}
}

// Register takes a RegisterProfileRequest, saves a Profile, and returns a RegisterProfileResponse.
func (service *RegistrationService) Register(request RegisterProfileRequest) (RegisterProfileResponse, error) {
	exists, err := service.repo.ExistsByUsername(request.Username)
	if err != nil {
		return RegisterProfileResponse{}, err
	}

	if exists {
		return RegisterProfileResponse{}, fmt.Errorf(ErrUsernameExists, request.Username)
	}

	profile := New(
		WithUsername(request.Username),
		WithPassword(request.Password),
	)

	err = service.repo.Save(profile)
	if err != nil {
		return RegisterProfileResponse{}, err
	}

	return RegisterProfileResponse{
		Id:       profile.id,
		Username: profile.username,
	}, nil
}
