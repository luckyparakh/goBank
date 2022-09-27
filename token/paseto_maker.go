package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto *paseto.V2
	symKey []byte
}

func NewPasetoMaker(key string) (Maker, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid secret key, it must of %v chars", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		symKey: []byte(key),
	}
	return maker, nil
}

func (p *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return p.paseto.Encrypt(p.symKey, payload, nil)
}
func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := p.paseto.Decrypt(token, p.symKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
