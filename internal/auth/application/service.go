package application

import (
	"time"

	"github.com/pkg/errors"
)

var (
	// ErrInvalidToken is when the token provided is not valid
	ErrInvalidToken = errors.New("invalid token provided")
	// ErrForbidden is when a user does not have the necessary scope to access a resource
	ErrForbidden = errors.New("resource forbidden")
)

// Account provided by an auth provider
type Account struct {
	// ID of the account e.g. email
	ID string `json:"id"`
	// Type of the account, e.g. service
	Type string `json:"type"`
	// Issuer of the account
	Issuer string `json:"issuer"`
	// Any other associated metadata
	Metadata map[string]string `json:"metadata"`
	// Scopes the account has access to
	Scopes []string `json:"scopes"`
	// Secret for the account, e.g. the password
	Secret string `json:"secret"`
}

// Token can be short or long lived
type Token struct {
	// The token to be used for accessing resources
	AccessToken string `json:"access_token"`
	// RefreshToken to be used to generate a new token
	RefreshToken string `json:"refresh_token"`
	// Time of token creation
	Created time.Time `json:"created"`
	// Time of token expiry
	Expiry time.Time `json:"expiry"`
}

// Expired returns a boolean indicating if the token needs to be refreshed
func (t *Token) Expired() bool {
	return t.Expiry.Unix() < time.Now().Unix()
}

type AuthService interface {
	// Init the auth
	Init(opts ...Option)
	// Options set for auth
	Options() Options
	// Inspect a token
	Inspect(token string) (*Account, error)
	// Token generated using refresh token or credentials
	Token(opts ...TokenOption) (*Token, error)
	// String returns the name of the implementation
	String() string
}

type authService struct {
	id string
}

func NewAuthService(id string) AuthService {
	return &authService{id: id}
}

func (a *authService) Init(opts ...Option) {
	//TODO implement me
	panic("implement me")
}

func (a *authService) Options() Options {
	//TODO implement me
	panic("implement me")
}

func (a *authService) Inspect(token string) (*Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authService) Token(opts ...TokenOption) (*Token, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authService) String() string {
	//TODO implement me
	panic("implement me")
}
