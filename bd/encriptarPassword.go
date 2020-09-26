package bd

import "golang.org/x/crypto/bcrypt"

/*encriptarPassword para seguridad de db */
func encriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
