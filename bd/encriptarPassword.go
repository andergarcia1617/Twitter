package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/* EncriptarPassword para seguridad */
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	byte, err := bcrypt.GenerateFromPassword([]byte(passs), costo)
	return string(bytes), err

}
