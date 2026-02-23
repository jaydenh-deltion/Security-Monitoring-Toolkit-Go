package systems

import (
	"fmt"
	"unicode"
)


func CheckPasswordStrength(password string) {
	// charset "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?"
	score := 0 

	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUppercase = true
		case char >= 'a' && char <= 'z':
			hasLowercase = true
		case char >= '0' && char <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
}

if hasUppercase { score += 25 }
if hasLowercase { score += 25 }
if hasDigit { score += 25 }
if hasSpecial { score += 25 }

fmt.Printf("Password: %s, Strength Score: %d/100\n", password, score)
}