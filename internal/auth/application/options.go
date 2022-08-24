package application

import (
	"time"
)

type Options struct {
	// Namespace the service belongs to
	Namespace string
	// ID is the services auth ID
	ID string
	// Secret is used to authenticate the service
	Secret string
	// Token is the services token used to authenticate itself
	Token *Token
	// PublicKey for decoding JWTs
	PublicKey string
	// PrivateKey for encoding JWTs
	PrivateKey string
	// Addrs sets the addresses of auth
	Addrs []string
}

func NewOptions(opts ...Option) Options {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	return options
}

type Option func(o *Options)

type TokenOptions struct {
	// ID for the account
	ID string
	// Secret for the account
	Secret string
	// RefreshToken is used to refesh a token
	RefreshToken string
	// Expiry is the time the token should live for
	Expiry time.Duration
}

type TokenOption func(o *TokenOptions)

// WithExpiry for the token
func WithExpiry(ex time.Duration) TokenOption {
	return func(o *TokenOptions) {
		o.Expiry = ex
	}
}

func WithCredentials(id, secret string) TokenOption {
	return func(o *TokenOptions) {
		o.ID = id
		o.Secret = secret
	}
}

func WithToken(rt string) TokenOption {
	return func(o *TokenOptions) {
		o.RefreshToken = rt
	}
}

// NewTokenOptions from a slice of options
func NewTokenOptions(opts ...TokenOption) TokenOptions {
	var options TokenOptions
	for _, o := range opts {
		o(&options)
	}

	// set defualt expiry of token
	if options.Expiry == 0 {
		options.Expiry = time.Minute
	}

	return options
}
