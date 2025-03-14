package types

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetObjectHandlerRequest struct {
	Task_id    string `json:"task_id"`
	Auth_token string `json:"auth_token"`
}

func CreateGetObjectHandlerRequest(r *http.Request) (*GetObjectHandlerRequest, error) {
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	task_id, okTask_id := req["task_id"]
	auth_token, okAuth_token := req["auth_token"]
	if !okAuth_token {
		return nil, fmt.Errorf("unauthorized")
	}
	if !okTask_id {
		return nil, fmt.Errorf("bad request")
	}
	return &GetObjectHandlerRequest{Task_id: task_id, Auth_token: auth_token}, nil
}

type GetObjectHandlerResultResponse struct {
	Result string `json:"result"`
}

type GetObjectHandlerStatusResponse struct {
	Status string `json:"status"`
}

type PostObjectHandlerAuthResponse struct {
	Auth_token string `json:"auth_token"`
}

type PostObjectHandlerTaskIdResponse struct {
	Task_id string `json:"task_id"`
}

type PostObjectHandlerTaskRequest struct {
	Task       string `json:"task"`
	Auth_token string `json:"auth_token"`
}

func CreatePostObjectHandlerRequest(r *http.Request) (*PostObjectHandlerTaskRequest, error) {
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	task, okTask := req["task"]
	auth_token, okAuth_token := req["auth_token"]
	if !okAuth_token {
		return nil, fmt.Errorf("unauthorized")
	}
	if !okTask {
		return nil, fmt.Errorf("bad request")
	}
	return &PostObjectHandlerTaskRequest{Task: task, Auth_token: auth_token}, nil
}

type PostObjectUserHandlerRequest struct {
	Login    string `json:"username"`
	Password string `json:"password"`
}

func CreatePostObjectHandlerUserRequest(r *http.Request) (*PostObjectUserHandlerRequest, error) {
	var req PostObjectUserHandlerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("error while decoding json: %v", err)
	}
	return &req, nil
}

func CreateObjectHandlerResponse(w http.ResponseWriter, err error, resp any) {
	if err != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}
