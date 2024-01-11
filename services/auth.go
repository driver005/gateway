package services

import (
	"context"

	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	ctx context.Context
	r   Registry
}

func NewAuthService(
	r Registry,
) *AuthService {
	return &AuthService{
		context.Background(),
		r,
	}
}

func (s *AuthService) SetContext(context context.Context) *AuthService {
	s.ctx = context
	return s
}

func (s *AuthService) ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) AuthenticateAPIToken(token string) types.AuthenticateResult {
	user, err := s.r.UserService().SetContext(s.ctx).RetrieveByApiToken(token, nil)
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
	passwordHash, err := s.r.UserService().SetContext(s.ctx).RetrieveByEmail(email, sql.Options{
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
		model, err := s.r.UserService().SetContext(s.ctx).RetrieveByEmail(email, sql.Options{})
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
	passwordHash, err := s.r.CustomerService().SetContext(s.ctx).RetrieveRegisteredByEmail(email, sql.Options{
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
			model, err := s.r.CustomerService().SetContext(s.ctx).RetrieveRegisteredByEmail(email, sql.Options{})
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
