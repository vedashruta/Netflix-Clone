package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Config Env
	Db     *mongo.Database
	Movies *mongo.Collection
	Users  *mongo.Collection
)

type Env struct {
	Server struct {
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"server"`
	Database struct {
		URI     string        `yaml:"uri"`
		Db      string        `yaml:"db"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"database"`
}

func Hash(input []byte) (hash string, err error) {
	hasher := md5.New()
	if _, err = hasher.Write(input); err != nil {
		return
	}
	hash = hex.EncodeToString(hasher.Sum(nil))
	return
}

func Encrypt(data []byte, passphrase string) ([]byte, error) {
	hash, err := Hash([]byte(passphrase))
	if err != nil {
		return nil, err
	}
	block, _ := aes.NewCipher([]byte(hash))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func Decrypt(data []byte, passphrase string) ([]byte, error) {
	hash, err := Hash([]byte(passphrase))
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(hash))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
