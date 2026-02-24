package systems

import (
	"fmt"
	"net/http"
)

// CheckSystemHealth checks the health of the system by making a simple HTTP request.
func CheckSystemHealth(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to reach the system: %v", err)
	}
	defer resp.Body.Close()