package generator

import (
	"crypto/rand"
	"math/big"
)

func Generate(length int) ([]byte, int, error) {
	b := make([]byte, length)
	n, err := rand.Read(b)
	return b, n, err
}

func MustGenerate(length int) []byte {
	b, _, _ := Generate(length)
	return b
}

func GenerateNumber(max *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}

func MustGenerateNumber(max *big.Int) *big.Int {
	n, _ := GenerateNumber(max)
	return n
}
