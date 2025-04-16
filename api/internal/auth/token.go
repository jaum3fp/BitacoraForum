package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jaum3fp/bitacora-forum/internal/models"
)

func GetPublicKey() *rsa.PublicKey {
	publicKey, err := loadPublicKeyFromEnv()
	if err != nil {
		log.Fatalf("Error obteniendo la clave pública: %v", err)
	}
	return publicKey
}

func GenerateUserToken(u models.User) (string, error) {

	privateKey, err := loadPrivateKeyFromEnv()
	if err != nil {
		return "", errors.New("authentication error")
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	t, err := tkn.SignedString(privateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", errors.New("authentication error")
	}

	return t, nil
}

func loadPrivateKeyFromEnv() (*rsa.PrivateKey, error) {
	privBase64 := os.Getenv("PRIVATE_KEY_BASE64")
	if privBase64 == "" {
		return nil, fmt.Errorf("clave privada no encontrada en las variables de entorno")
	}

	privData, err := base64.StdEncoding.DecodeString(privBase64)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la clave privada: %v", err)
	}

	block, _ := pem.Decode(privData)
	if block == nil {
		return nil, fmt.Errorf("clave privada no válida")
	}

	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error al analizar la clave privada: %v", err)
	}

	return privKey.(*rsa.PrivateKey), nil
}

func loadPublicKeyFromEnv() (*rsa.PublicKey, error) {
	pubBase64 := os.Getenv("PUBLIC_KEY_BASE64")
	if pubBase64 == "" {
		return nil, fmt.Errorf("no se ha encuentrado la clave pública")
	}

	pubData, err := base64.StdEncoding.DecodeString(pubBase64)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la clave pública: %v", err)
	}

	block, _ := pem.Decode(pubData)
	if block == nil {
		return nil, fmt.Errorf("clave pública no válida")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error al analizar la clave pública: %v", err)
	}

	return pubKey.(*rsa.PublicKey), nil
}
