package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ZolotarevAlexandr/yl_sprint_1_final/calculator/calculator"
)

type RequestPayload struct {
	Expression string `json:"expression"`
}

type ResponsePayload struct {
	Result *float64 `json:"result,omitempty"`
	Error  *string  `json:"error,omitempty"`
}

var calculationErrors = map[string]int{
	"not enough operands":     http.StatusUnprocessableEntity,
	"invalid stack":           http.StatusUnprocessableEntity,
	"invalid operand":         http.StatusUnprocessableEntity,
	"mismatched parentheses":  http.StatusUnprocessableEntity,
	"unsupported token value": http.StatusUnprocessableEntity,
	"division by zero":        http.StatusUnprocessableEntity,
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var requestPayload RequestPayload
	err = json.Unmarshal(body, &requestPayload)
	if err != nil {
		errorResponse(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	result, err := calculator.Calculate(requestPayload.Expression)
	if err != nil {
		if status, exists := calculationErrors[err.Error()]; exists {
			errorResponse(w, "Expression is not valid", status)
			return
		}
		errorResponse(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := ResponsePayload{
		Result: &result,
	}
	jsonResponse(w, response, http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, response ResponsePayload, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func errorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := ResponsePayload{
		Error: &message,
	}
	jsonResponse(w, response, statusCode)
}
