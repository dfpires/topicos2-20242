package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Função para gerar o hash da senha
func GeneratePasswordHash(password string) (string, error) {
	// Cria o hash da senha com custo 14 (custo padrão é 10)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Função para verificar a senha
func CheckPasswordHash(password, hash string) bool {
	// Compara a senha com o hash armazenado
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "minhasenha123"

	// Gerando o hash da senha
	hash, err := GeneratePasswordHash(password)
	if err != nil {
		fmt.Println("Erro ao gerar o hash da senha:", err)
		return
	}

	fmt.Println("Hash gerado:", hash)

	// Verificando a senha
	match := CheckPasswordHash("minhasenha1234", hash)
	if match {
		fmt.Println("Senha correta!")
	} else {
		fmt.Println("Senha incorreta!")
	}
}
