package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ContentType       = "Content-Type"
	ContentTypeByJSON = "application/json"
	ContentTypeByText = "text/plain"
)

// 读取请求体，映射为实体
func ParseJSON[T any](r *http.Request) (*T, error) {
	if r.Body == nil {
		return nil, fmt.Errorf("missing request body")
	}
	req := new(T)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}
	return req, nil
}

// Response OK
func Ok(w http.ResponseWriter, value any) error {
	return WriteJSON(w, http.StatusOK, value)
}

// Response Fail
func Fail(w http.ResponseWriter, message string) error {
	resp := map[string]any{"message": message}
	return WriteJSON(w, http.StatusInternalServerError, resp)
}

// Response JSON
func WriteJSON(w http.ResponseWriter, httpStatus int, value any) error {
	w.Header().Add(ContentType, ContentTypeByJSON)
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(value)
}
