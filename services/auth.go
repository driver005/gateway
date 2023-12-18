package services

import (
	"github.com/driver005/gateway/repository"
	"github.com/driver005/gateway/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService     UserService
	customerService CustomerService
}

func NewAuthService(
	userService UserService,
	customerService CustomerService,
) *AuthService {
	return &AuthService{
		userService,
		customerService,
	}
}

func (s *AuthService) ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) AuthenticateAPIToken(token string) types.AuthenticateResult {
	user, err := s.userService.RetrieveByApiToken(token, nil)
	if err != nil {
		return types.AuthenticateResult{
			Error:   "Invalid API Token",
			Success: false,
		}
	}

	return types.AuthenticateResult{
		Success: true,
		User:    user,
	}
}

func (s *AuthService) Authenticate(email string, password string) types.AuthenticateResult {
	passwordHash, err := s.userService.RetrieveByEmail(email, repository.Options{
		Selects: []string{"password_hash"},
	})
	if err != nil {
		return types.AuthenticateResult{
			Error:   "Invalid email or password",
			Success: false,
		}
	}

	passwordsMatch := s.ComparePassword(password, passwordHash.PasswordHash)
	if passwordsMatch {
		model, err := s.userService.RetrieveByEmail(email, repository.Options{})
		if err != nil {
			return types.AuthenticateResult{
				Error:   "Invalid email or password",
				Success: false,
			}
		}

		return types.AuthenticateResult{
			Success: true,
			User:    model,
		}
	}

	return types.AuthenticateResult{
		Error:   "Invalid email or password",
		Success: false,
	}
}

func (s *AuthService) AuthenticateCustomer(email string, password string) types.AuthenticateResult {
	passwordHash, err := s.customerService.RetrieveRegisteredByEmail(email, repository.Options{
		Selects: []string{"password_hash"},
	})
	if err != nil {
		return types.AuthenticateResult{
			Error:   "Invalid email or password",
			Success: false,
		}
	}

	if passwordHash.PasswordHash != "" {
		passwordsMatch := s.ComparePassword(password, passwordHash.PasswordHash)
		if passwordsMatch {
			model, err := s.customerService.RetrieveRegisteredByEmail(email, repository.Options{})
			if err != nil {
				return types.AuthenticateResult{
					Error:   "Invalid email or password",
					Success: false,
				}
			}

			return types.AuthenticateResult{
				Success:  true,
				Customer: model,
			}
		}
	}

	return types.AuthenticateResult{
		Error:   "Invalid email or password",
		Success: false,
	}
}
