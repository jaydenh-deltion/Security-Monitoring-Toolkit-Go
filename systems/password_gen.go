package systems

import (
	"fmt"
	"math/rand"
	"time"
)

type PasswordGenerator struct {
	length int
}

func (pg *PasswordGenerator) Generate() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	password := make([]byte, pg.length)
	
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

func NewPasswordGenerator(length int) *PasswordGenerator {
	rand.Seed(time.Now().UnixNano())
	return &PasswordGenerator{length: length}
}

func main() {
	pg := NewPasswordGenerator(12)
	fmt.Println("Generated Password:", pg.Generate())
}