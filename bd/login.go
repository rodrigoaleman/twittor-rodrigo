package bd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"twittor-rodrigo/models"
)

func Login(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		log.Fatal("Contrasena incorrecta")
		return usu, false
	}
	return usu, true
}
