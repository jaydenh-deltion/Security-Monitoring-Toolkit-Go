package systems

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type PasswordScore struct {
		score int

	}

func RunChecker() {
	fmt.Println("Checking password strength...")
	fmt.Println("Please enter your password:")
	var password string
	fmt.Scanln(&password)

	strength := CheckPasswordStrength(password)
	fmt.Printf("Password strength: %s\n", strength)
}

func CheckPasswordStrength(password string) string {
	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case !unicode.IsLetter(char) && !unicode.IsDigit(char):
			hasSpecial = true
		}
	}
	length := utf8.RuneCountInString(password)

		score := 0
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasDigit {
		score++
	}
	if hasSpecial {
		score++
	}
	if length >= 12 {
		score++
	}
	
	switch score {
	case 5:
		return "Excellent"
	case 4:
		return "Strong"
	case 3:
		return "Moderate"
	case 2:
		return "Weak"
	default:
		return "Very Weak"
	}
}