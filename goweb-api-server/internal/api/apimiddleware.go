package api

import (
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	requests     = make(map[string][]time.Time)
	mutex        = &sync.Mutex{}
	requestLimit = 100 // Number of allowed requests
	windowSize   = time.Minute
)

// APIKeyMiddleware checks the Authorization header for a valid API key
func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the API key from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid API key", http.StatusUnauthorized)
			return
		}

		// Extract the API key from the header
		apiKey := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the API key (this could be a check against a database or in-memory store)
		if apiKey != "your-secret-api-key" {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		// Rate limiting
		if !allowRequest(apiKey) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		// If the API key is valid and within the rate limit, call the next handler
		next.ServeHTTP(w, r)
	})
}

func allowRequest(apiKey string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now()
	if requestTimes, exists := requests[apiKey]; exists {
		// Filter out requests that are outside the window
		filteredRequests := []time.Time{}
		for _, t := range requestTimes {
			if now.Sub(t) <= windowSize {
				filteredRequests = append(filteredRequests, t)
			}
		}

		// Update the request history
		requests[apiKey] = append(filteredRequests, now)

		// Check if the request count exceeds the limit
		return len(requests[apiKey]) <= requestLimit
	}

	// If no previous requests, initialize the slice
	requests[apiKey] = []time.Time{now}
	return true
}
