// Package handlers wires HTTP requests to the calculator package and
// handles JSON serialization, validation and error mapping
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SharonBarrial/calculator-backend/calculator"
)

// writeJSON is a small helper to avoid reating
// "set header, write status, encode JSON" boilerplate in every handler
func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

// decodeBinaryRequest parses and validates a JSON request body for binary operations (add, subtract, multiply, divide, power)
// it returns false if the body is malformed or if the fields are not numeric, and writes a 400 Bad Request response
func decodeBinaryRequest(w http.ResponseWriter, r *http.Request) (BinaryOperationRequest, bool) {
	var req BinaryOperationRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body: expected JSON with numeric fields \"a\" and \"b\"")
		return req, false
	}
	return req, true
}

func decodeUnaryRequest(w http.ResponseWriter, r *http.Request) (UnaryOperationRequest, bool) {
	var req UnaryOperationRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body: expected JSON with numeric field \"a\"")
		return req, false
	}
	return req, true
}

// Add handles POST /api/add
func Add(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, OperationResponse{Result: calculator.Add(req.A, req.B)})
}

// Subtract handles POST /api/subtract
func Subtract(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, OperationResponse{Result: calculator.Subtract(req.A, req.B)})
}

// Multiply handles POST /api/multiply
func Multiply(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, OperationResponse{Result: calculator.Multiply(req.A, req.B)})
}

// Divide handles POST /api/divide
func Divide(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}
	result, err := calculator.Divide(req.A, req.B)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, OperationResponse{Result: result})
}

// Power handles POST /api/power
func Power(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, OperationResponse{Result: calculator.Power(req.A, req.B)})
}

// Sqrt handles POST /api/sqrt
func Sqrt(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeUnaryRequest(w, r)
	if !ok {
		return
	}
	result, err := calculator.Sqrt(req.A)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, OperationResponse{Result: result})
}

// Percentage handles POST /api/percentage
// Interprest the request as; "what percentage is A of B"
func Percentage(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeBinaryRequest(w, r)
	if !ok {
		return
	}
	result, err := calculator.Percentage(req.A, req.B)
	if err != nil {
		if errors.Is(err, calculator.ErrInvalidPercentage) {
			writeError(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}
	writeJSON(w, http.StatusOK, OperationResponse{Result: result})
}

// Health handles GET /api/health — useful for Docker healthchecks and
// for the frontend to check backend availability
func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
