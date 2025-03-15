package executor

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ExecuteRequest sends the formatted API request and returns the response as a string.
// Precondition: request is a valid API request string formatted as method URL.
// Postcondition: returns the API response or an error message.
func ExecuteRequest(request string, headers map[string]string, body string) string {
	// Split the request string into method and URL
	parts := strings.SplitN(request, " ", 2)
	if len(parts) < 2 {
		return "Invalid request format. Expected: 'METHOD URL'."
	}

	method := strings.ToUpper(parts[0])
	url := parts[1]

	// Create the request body as a reader if there is a body
	var bodyReader io.Reader
	if body != "" {
		bodyReader = bytes.NewBuffer([]byte(body))
	}

	// Create the HTTP request
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Sprintf("Error creating request: %v", err)
	}

	// Set the headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform the request using the http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error executing request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v", err)
	}

	// Return the response as a string
	return string(responseBody)
}
