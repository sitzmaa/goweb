package receiver

import "fmt"

// HandleResponse processes the API response.
// Precondition: response is a non-empty string.
// Postcondition: prints the response to the console.
func HandleResponse(response string) {
	fmt.Println("API Response:", response)
}
