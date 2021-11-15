package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

//Similar to what we’ve done with JWT,we declare type PasetoMaker struct,
//which will implement the same token.Maker interface, but use PASETO instead of JWT.
type PasetoMaker struct {
    paseto       *paseto.V2
    symmetricKey []byte
}


func NewPasetoMaker(symmetricKey string) (Maker, error) {
	//Paseto version 2 uses Chacha20 Poly1305 algorithm to encrypt the payload. 
	//So here we have to check the length of the symmetric key
	// to make sure that it has the correct size that’s required by the algorithm
    if len(symmetricKey) != chacha20poly1305.KeySize {
        return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
    }

	//Else, we just create a new PasetoMaker object that contains paseto.NewV2() and the input symmetricKey converted to []byte slice.
    maker := &PasetoMaker{
        paseto:       paseto.NewV2(),
        symmetricKey: []byte(symmetricKey),
    }

    return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
    payload, err := NewPayload(username, duration)
    if err != nil {
        return "", err
    }

    return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
    payload := &Payload{}

    err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
    if err != nil {
        return nil, ErrInvalidToken
    }

	//we will check if the token is valid or not by calling payload.Valid().
    err = payload.Valid()
    if err != nil {
        return nil, err
    }

    return payload, nil
}