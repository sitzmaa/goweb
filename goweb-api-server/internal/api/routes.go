package api

import (
	"encoding/json"
	"fmt"
	"goweb-server/internal/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Define RESTful endpoints and map them to handler functions
	r.HandleFunc("/api/get", getHandler).Methods("GET")
	r.HandleFunc("/api/post", postHandler).Methods("POST")
	r.HandleFunc("/api/put", putHandler).Methods("PUT")
	r.HandleFunc("/api/patch", patchHandler).Methods("PATCH")
	r.HandleFunc("/api/delete", deleteHandler).Methods("DELETE")
	r.HandleFunc("/api/v1/deploy", handleDeployment).Methods("POST")
	r.HandleFunc("/api/v1/log", handleLog).Methods("POST")
	r.HandleFunc("/api/v1/persist", handleDatabaseWrite).Methods("POST")
}

func handleRequest(w http.ResponseWriter, r *http.Request, req models.Request) {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		handleError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		handleError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request validated successfully"))
}

func handleDeployment(w http.ResponseWriter, r *http.Request) {
	http.Post("http://localhost:5002/deploy", "application/json", nil)
	fmt.Fprintln(w, "Service Deployed Successfully")
}

func handleLog(w http.ResponseWriter, r *http.Request) {
	http.Post("http://localhost:5003/log", "application/json", nil)
	fmt.Fprintln(w, "Log Event Sent")
}

func handleDatabaseWrite(w http.ResponseWriter, r *http.Request) {
	http.Post("http://localhost:5001/write", "application/json", nil)
	fmt.Fprintln(w, "Data Persisted")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.GetRequest{}
	handleRequest(w, r, req)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.PostRequest{}
	handleRequest(w, r, req)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.PutRequest{}
	handleRequest(w, r, req)
}

func patchHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.PatchRequest{}
	handleRequest(w, r, req)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.DeleteRequest{}
	handleRequest(w, r, req)
}

func handleError(w http.ResponseWriter, message string, statusCode int) {
	log.Printf("Error: %s", message)
	http.Error(w, message, statusCode)
}
