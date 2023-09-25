package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
// 	return payload.parseClaimString("ID")
// }

// func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
// 	return jwt.NewNumericDate(payload.ExpiredAt), nil
// }

// func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
// 	return jwt.NewNumericDate(payload.IssuedAt), nil
// }

// func (payload *Payload) GetIssuer() (string, error) {
// 	return "", nil
// }
