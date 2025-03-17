package actions

import (
	"fmt"
	"goweb-server/internal/database"
	"io"
	"log"
	"net/http"
)

// Perform GET action
func PerformGetAction() string {
	// Here you would define the logic for handling the GET request
	// For simplicity, let's assume it returns a dummy response
	return "GET request processed successfully!"
}

// Perform POST action
func PerformPostAction(r *http.Request) string {
	// Read the body of the POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Sprintf("Error reading POST body: %v", err)
	}
	// Process the body and return a response
	return fmt.Sprintf("POST request received with data: %s", string(body))
}

// Perform DELETE action
func PerformDeleteAction(r *http.Request) string {
	// You can handle the DELETE request here
	// For now, let's return a simple response
	return "DELETE request processed successfully!"
}

// Perform PUT action
func PerformPutAction(r *http.Request) string {
	// Read the body of the PUT request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Sprintf("Error reading PUT body: %v", err)
	}
	// Process the body and return a response
	return fmt.Sprintf("PUT request received with data: %s", string(body))
}

// LogAction stores a log entry in the database
func LogAction(action, status string) {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO logs (action, status) VALUES (?, ?)", action, status)
	if err != nil {
		log.Println("Error inserting log:", err)
	}
}
