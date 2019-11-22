package http

// Payload must be a json serializable structure. Payload
// represents the application specific message you are transporting
type Payload interface{}

// Response is a standard json serializable response schema.
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	// Additional must be json serializable or expect errors
	Payload `json:"payload,omitempty"`
}

func NewResponse(code string, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

func NewResponseWithPayload(code string, message string, payload interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Payload: payload,
	}
}
