package executor

import (
	"fmt"
	"io"
	"net/http"
)

// ExecuteRequest sends the formatted API request and returns the response as a string.
// Precondition: request is a valid API request string.
// Postcondition: returns the API response or an error message.
func ExecuteRequest(request string) string {
	resp, err := http.Get(request)
	if err != nil {
		return fmt.Sprintf("Error executing request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v", err)
	}

	return string(body)
}
