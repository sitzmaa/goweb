package interpreter

import (
	"fmt"
	"goweb/internal/executor"
	"goweb/internal/receiver"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Command struct {
	APIRequests []struct {
		Method      string `yaml:"method"`
		URL         string `yaml:"url"`
		Description string `yaml:"description"`
		Headers     struct {
			ContentType string `yaml:"Content-Type"`
		} `yaml:"headers"`
		Params struct {
			Page    int `yaml:"page"`
			PerPage int `yaml:"per_page"`
		} `yaml:"params,omitempty"`
		Response struct {
			Status int `yaml:"status"`
			Body   []struct {
				ID          int    `yaml:"id"`
				Name        string `yaml:"name"`
				Description string `yaml:"description"`
			} `yaml:"body"`
		} `yaml:"response"`
		Body struct {
			Name        string `yaml:"name"`
			Description string `yaml:"description"`
		} `yaml:"body,omitempty"`
	} `yaml:"api_requests"`
}

// HandleCommand processes the command input from the user.
// It directly parses the argument string and constructs the API call accordingly.
func HandleCommand(command string) {
	// Split the command into parts (e.g., "goweb", "get", "--ip=192.168.1.1", "--port=8080")
	parts := strings.Fields(command)

	if len(parts) < 2 {
		log.Println("Invalid command. Please specify the action and necessary parameters.")
		return
	}

	// Extract the action (GET, POST, DELETE, SEND)
	action := parts[1]
	switch action {
	case "get":
		handleGet(parts)
	case "post":
		handlePost(parts)
	case "delete":
		handleDelete(parts)
	case "send":
		handleSend(parts)
	default:
		log.Printf("Unknown action: %s. Supported actions are: get, post, delete, send.\n", action)
	}
}

// handleGet processes the GET command and builds the API request.
func handleGet(parts []string) {
	ip := ""
	port := ""

	// Parse arguments
	for _, part := range parts[2:] {
		if strings.HasPrefix(part, "--ip=") {
			ip = strings.TrimPrefix(part, "--ip=")
		} else if strings.HasPrefix(part, "--port=") {
			port = strings.TrimPrefix(part, "--port=")
		}
	}

	if ip == "" || port == "" {
		log.Println("Missing --ip or --port arguments for GET request.")
		return
	}

	// Construct the GET request URL
	url := fmt.Sprintf("http://%s:%s", ip, port)
	fmt.Println("GET Request URL:", url)

	// Create headers (empty in this case)
	headers := make(map[string]string)

	// Execute the GET request
	response := executor.ExecuteRequest("GET", headers, "")
	receiver.HandleResponse(response)
}

// handlePost processes the POST command and builds the API request.
func handlePost(parts []string) {
	url := ""
	data := ""
	headers := make(map[string]string)

	// Parse arguments
	for _, part := range parts[2:] {
		if strings.HasPrefix(part, "--url=") {
			url = strings.TrimPrefix(part, "--url=")
		} else if strings.HasPrefix(part, "--data=") {
			data = strings.TrimPrefix(part, "--data=")
		}
	}

	if url == "" || data == "" {
		log.Println("Missing --url or --data arguments for POST request.")
		return
	}

	// Set headers for POST request
	headers["Content-Type"] = "application/json"

	// Execute the POST request
	response := executor.ExecuteRequest("POST", headers, data)
	receiver.HandleResponse(response)
}

// handleDelete processes the DELETE command and builds the API request.
func handleDelete(parts []string) {
	url := ""
	headers := make(map[string]string)

	// Parse arguments
	for _, part := range parts[2:] {
		if strings.HasPrefix(part, "--url=") {
			url = strings.TrimPrefix(part, "--url=")
		}
	}

	if url == "" {
		log.Println("Missing --url argument for DELETE request.")
		return
	}

	// Execute the DELETE request
	response := executor.ExecuteRequest("DELETE", headers, "")
	receiver.HandleResponse(response)
}

// handleSend processes the SEND command and calls HandleYaml to process the YAML file.
func handleSend(parts []string) {
	if len(parts) < 3 {
		log.Println("Missing YAML file argument for SEND command.")
		return
	}

	filePath := parts[2]
	HandleYaml(filePath)
}

// HandleYaml processes the YAML file input from the user.
// Precondition: The file path is a valid string input.
// Postcondition: The YAML file is parsed, executed, and the response is handled.
func HandleYaml(filePath string) {
	var command Command
	// Read the YAML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	// Parse the YAML file into the Command struct
	err = yaml.Unmarshal(data, &command)
	if err != nil {
		log.Fatalf("failed to read yaml: %s", err)
	}

	// Loop through each API request and generate an HTTP request string
	for _, req := range command.APIRequests {
		// Prepare the full URL with query parameters for GET requests
		fullURL := req.URL
		if req.Method == "GET" && req.Params.Page != 0 && req.Params.PerPage != 0 {
			fullURL = fmt.Sprintf("%s?page=%d&per_page=%d", req.URL, req.Params.Page, req.Params.PerPage)
		}

		// Initialize the request body as an empty string
		var reqBody string
		if req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH" {
			// If there is a body (for POST, PUT, or PATCH), create a JSON string
			reqBody = fmt.Sprintf(`{"name": "%s", "description": "%s"}`, req.Body.Name, req.Body.Description)
		}

		// Create headers as a map to pass to ExecuteRequest
		headers := map[string]string{
			"Content-Type": req.Headers.ContentType,
		}

		// Call the ExecuteRequest function
		response := executor.ExecuteRequest(req.Method+" "+fullURL, headers, reqBody)

		// Print the response for debugging
		fmt.Printf("Response for %s %s:\n%s\n", req.Method, fullURL, response)

		// Handle the response, pass it to receiver
		receiver.HandleResponse(response)
	}
}
