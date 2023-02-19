package token

import "time"

type Maker interface {
	// GenerateToken Generating token from username and duration
	GenerateToken(username int64, email string, duration time.Duration) (string, error)

	// VerifyToken will verifying does the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
