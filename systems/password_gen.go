package systems

import (
	"fmt"
	"crypto/rand"
	"math/big"
)

type PasswordGenerator struct {
	length int
}

func (pg *PasswordGenerator) Generate() (string, error){
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?"
	password := make([]byte, pg.length)

	max := big.NewInt(int64(len(charset)))

	for i := range password {
		nBig, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		password[i] = charset[nBig.Int64()]
	}
	
	return string(password), nil
}

func NewPasswordGenerator(length int) *PasswordGenerator {
	return &PasswordGenerator{length: length}
}

func main() {

	var length int 
	fmt.Print("Enter desired password length: ")
	fmt.Scanln(&length)

	if length < 8 {
		var confirm string
		fmt.Printf("WARNING: Password length must be at least 8 characters. a password of %d characters may be easily guessable.", length)
		fmt.Scanln(&confirm)

		if confirm != "y"  {
			fmt.Println("Aborting password generation.")
			return
		}
	}

	pg := NewPasswordGenerator(length)
	password, err := pg.Generate()
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}
	fmt.Println("Generated password:", password)

	// Example of generating a random number using crypto/rand
	nBig, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}
	fmt.Println("Random number:", nBig.Int64())
}