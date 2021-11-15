package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type Payload struct {
	ID        uuid.UUID `json:"id"` //uniquely identify each token
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"` //when the token is created
	ExpiredAt time.Time `json:"expired_at"`
}


func NewPayload(username string, duration time.Duration) (*Payload, error) {
    tokenID, err := uuid.NewRandom() //generate a unique token ID
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


var (
    ErrInvalidToken = errors.New("token is invalid")
    ErrExpiredToken = errors.New("token has expired")
)
func (payload *Payload) Valid() error {
    if time.Now().After(payload.ExpiredAt) {
		//We should declare this error as a public constant: 
        return ErrExpiredToken
    }
    return nil
}