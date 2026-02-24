package systems

import (
	"fmt"
	"unicode"
)

// CheckPassword checks if the password meets the specified criteria.
func CheckPassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}