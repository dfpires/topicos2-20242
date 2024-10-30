package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Chave secreta para assinar o token
var jwtKey = []byte("minha_chave_secreta")

// Estrutura de payload do JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Função para gerar um token JWT
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Função para verificar o token JWT
func VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, fmt.Errorf("token expirado")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	return claims, nil
}

func main() {
	username := "usuario_teste"

	// Gera o token JWT
	token, err := GenerateJWT(username)
	if err != nil {
		fmt.Println("Erro ao gerar token:", err)
		return
	}

	fmt.Println("Token gerado:", token)

	// Verifica o token JWT
	claims, err := VerifyJWT(token)
	if err != nil {
		fmt.Println("Erro ao verificar token:", err)
		return
	}

	fmt.Printf("Token verificado! Usuário: %s\n", claims.Username)
}
