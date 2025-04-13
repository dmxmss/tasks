package repository

import (
	"github.com/dmxmss/tasks/config"
	e "github.com/dmxmss/tasks/error"
	"golang.org/x/crypto/bcrypt"
)

type HashRepository interface {
	HashPassword(string) (string, error)
	IsPasswordValid(string, string) bool
}

type hashRepository struct {
	conf *config.Hash
}

func NewHashRepository(conf *config.Hash) HashRepository {
	return &hashRepository{
		conf: conf,
	}
}

func (h *hashRepository) HashPassword(raw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(raw), h.conf.Cost)
	if err != nil {
		return "", e.ErrHashingFailed
	}
	return string(bytes), err
}

func (h *hashRepository) IsPasswordValid(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
