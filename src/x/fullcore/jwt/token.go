package jwt

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Maker interface {
	VerifyToken(token string) (*Payload, error)
}

type JWTMaker struct {
	secretKey string
}
