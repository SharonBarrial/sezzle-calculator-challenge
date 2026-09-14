package handlers

//BinaryOperationRequest is the JSON body expected by
//the /calculate endpoint for binary operations
//like add, subtract, multiply, divide, and power
type BinaryOperationRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

//UnaryOperationRequest is the JSON body expected
//by one-operand operantions like square root
type UnaryOperationRequest struct {
	A float64 `json:"a"`
}

//OperationResponse is returned on successful calculation
type OperationResponse struct {
	Result float64 `json:"result"`
}

//ErrorResponse is returned on any 4xx or 5xx error from the API
type ErrorResponse struct {
	Error string `json:"error"`
}
