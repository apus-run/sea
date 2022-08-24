package auth

import (
	"github.com/apus-run/sea/internal/auth/application"
)

// Module is a struct that defines all dependencies inside hash module
type Module struct {
	authService application.AuthService
}

// Configure setups all dependencies
func (m *Module) Configure(id string) {
	m.authService = application.NewAuthService(id)
}

// GetAuthService returns an instance of application.AuthService
func (m Module) GetAuthService() application.AuthService {
	return m.authService
}
