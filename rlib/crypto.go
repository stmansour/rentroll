package rlib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Encrypt a slice of bytes using the server's key
// Reference: https://play.golang.org/p/mpXKSF9fdC9
//
// INPUTS:
//     s - string to encrypt
//
// RETURNS:
//     encrypted slice of bytes
//     any error encountered
//-----------------------------------------------------------------------------
func Encrypt(s string) ([]byte, error) {
	return EncryptCore([]byte(s), RRdb.Key)
}

// EncryptCore a slice of bytes using the server's key
// Reference: https://play.golang.org/p/mpXKSF9fdC9
//
// INPUTS:
//     p   - array of bytes to encrypt
//     key - crypto key
//
// RETURNS:
//     encrypted slice of bytes
//     any error encountered
//-----------------------------------------------------------------------------
func EncryptCore(p []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, p, nil), nil
}

// Decrypt a slice of bytes using the server's key
// Reference: https://play.golang.org/p/mpXKSF9fdC9
//
// INPUTS:
//     b - slice of bytes to decript
//
// RETURNS:
//     decrypted string
//     any error encountered
//-----------------------------------------------------------------------------
func Decrypt(b []byte) (string, error) {
	d, err := DecryptCore(b, RRdb.Key)
	return string(d), err
}

// DecryptCore decripts a slice of bytes using the supplied key
// Reference: https://play.golang.org/p/mpXKSF9fdC9
//
// INPUTS:
//     p   - array of bytes to encrypt
//     key - crypto key
//
//     decrypted slice of bytes
//     any error encountered
//-----------------------------------------------------------------------------
func DecryptCore(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
