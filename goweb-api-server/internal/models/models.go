package models

import "fmt"

// Request interface for validation
type Request interface {
	Validate() error
}

type PostRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (r *PostRequest) Validate() error {
	if r.Username == "" {
		return fmt.Errorf("username is required")
	}
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

type PutRequest struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

func (r *PutRequest) Validate() error {
	if r.ID == "" {
		return fmt.Errorf("ID is required")
	}
	if r.Data == "" {
		return fmt.Errorf("data is required")
	}
	return nil
}

type PatchRequest struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func (r *PatchRequest) Validate() error {
	if r.Field == "" {
		return fmt.Errorf("field is required")
	}
	if r.Value == "" {
		return fmt.Errorf("value is required")
	}
	return nil
}

type DeleteRequest struct {
	ID string `json:"id"`
}

func (r *DeleteRequest) Validate() error {
	if r.ID == "" {
		return fmt.Errorf("ID is required")
	}
	return nil
}

type GetRequest struct {
	Query string `json:"query"`
}

func (r *GetRequest) Validate() error {
	if r.Query == "" {
		return fmt.Errorf("query is required")
	}
	return nil
}
