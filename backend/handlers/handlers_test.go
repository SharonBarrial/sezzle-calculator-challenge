package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// doRequest builds an httptest request/recorder pair, invokes the handler,
// and returns the recorder so tests can inspect status code and body.
func doRequest(t *testing.T, handler http.HandlerFunc, body any) *httptest.ResponseRecorder {
	t.Helper()

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("failed to encode request body: %v", err)
		}
	}

	req := httptest.NewRequest(http.MethodPost, "/api/test", &buf)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler(rec, req)
	return rec
}

func TestAddHandler(t *testing.T) {
	rec := doRequest(t, Add, BinaryOperationRequest{A: 2, B: 3})

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d (body: %s)", rec.Code, rec.Body.String())
	}

	var resp OperationResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.Result != 5 {
		t.Errorf("expected result 5, got %v", resp.Result)
	}
}

func TestSubtractHandler(t *testing.T) {
	rec := doRequest(t, Subtract, BinaryOperationRequest{A: 10, B: 4})

	var resp OperationResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if rec.Code != http.StatusOK || resp.Result != 6 {
		t.Errorf("expected 200/6, got %d/%v", rec.Code, resp.Result)
	}
}

func TestMultiplyHandler(t *testing.T) {
	rec := doRequest(t, Multiply, BinaryOperationRequest{A: 6, B: 7})

	var resp OperationResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if rec.Code != http.StatusOK || resp.Result != 42 {
		t.Errorf("expected 200/42, got %d/%v", rec.Code, resp.Result)
	}
}

func TestDivideHandler(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		rec := doRequest(t, Divide, BinaryOperationRequest{A: 10, B: 2})

		var resp OperationResponse
		json.Unmarshal(rec.Body.Bytes(), &resp)

		if rec.Code != http.StatusOK || resp.Result != 5 {
			t.Errorf("expected 200/5, got %d/%v", rec.Code, resp.Result)
		}
	})

	t.Run("division by zero returns 422", func(t *testing.T) {
		rec := doRequest(t, Divide, BinaryOperationRequest{A: 10, B: 0})

		if rec.Code != http.StatusUnprocessableEntity {
			t.Fatalf("expected status 422, got %d", rec.Code)
		}

		var resp ErrorResponse
		json.Unmarshal(rec.Body.Bytes(), &resp)
		if resp.Error == "" {
			t.Error("expected a non-empty error message")
		}
	})
}

func TestPowerHandler(t *testing.T) {
	rec := doRequest(t, Power, BinaryOperationRequest{A: 2, B: 10})

	var resp OperationResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if rec.Code != http.StatusOK || resp.Result != 1024 {
		t.Errorf("expected 200/1024, got %d/%v", rec.Code, resp.Result)
	}
}

func TestSqrtHandler(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		rec := doRequest(t, Sqrt, UnaryOperationRequest{A: 16})

		var resp OperationResponse
		json.Unmarshal(rec.Body.Bytes(), &resp)

		if rec.Code != http.StatusOK || resp.Result != 4 {
			t.Errorf("expected 200/4, got %d/%v", rec.Code, resp.Result)
		}
	})

	t.Run("negative number returns 422", func(t *testing.T) {
		rec := doRequest(t, Sqrt, UnaryOperationRequest{A: -9})

		if rec.Code != http.StatusUnprocessableEntity {
			t.Fatalf("expected status 422, got %d", rec.Code)
		}
	})
}

func TestPercentageHandler(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		rec := doRequest(t, Percentage, BinaryOperationRequest{A: 50, B: 200})

		var resp OperationResponse
		json.Unmarshal(rec.Body.Bytes(), &resp)

		if rec.Code != http.StatusOK || resp.Result != 25 {
			t.Errorf("expected 200/25, got %d/%v", rec.Code, resp.Result)
		}
	})

	t.Run("zero base returns 422", func(t *testing.T) {
		rec := doRequest(t, Percentage, BinaryOperationRequest{A: 50, B: 0})

		if rec.Code != http.StatusUnprocessableEntity {
			t.Fatalf("expected status 422, got %d", rec.Code)
		}
	})
}

func TestMalformedBodyReturns400(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBufferString(`{"a": "not-a-number"}`))
	rec := httptest.NewRecorder()

	Add(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestUnknownFieldReturns400(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/add", bytes.NewBufferString(`{"a": 1, "b": 2, "c": 3}`))
	rec := httptest.NewRecorder()

	Add(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for unknown field, got %d", rec.Code)
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()

	Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}
