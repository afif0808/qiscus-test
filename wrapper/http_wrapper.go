package wrapper

import (
	"encoding/json"
	"net/http"

	"github.com/afif0808/qiscus-test/meta"
)

// Meta model

// HTTPResponse format
type HTTPResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// NewHTTPResponse for create common response
func NewHTTPResponse(code int, message string, params ...interface{}) *HTTPResponse {
	commonResponse := new(HTTPResponse)

	for _, param := range params {
		switch val := param.(type) {
		case *meta.Meta, meta.Meta:
			commonResponse.Meta = val
		case error:
			commonResponse.Error = val.Error()
		default:
			commonResponse.Data = param
		}
	}

	if code < http.StatusBadRequest {
		commonResponse.Success = true
	}
	commonResponse.Code = code
	commonResponse.Message = message
	return commonResponse
}

// JSON for set http JSON response (Content-Type: application/json) with parameter is http response writer
func (resp *HTTPResponse) JSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	return json.NewEncoder(w).Encode(resp)
}
