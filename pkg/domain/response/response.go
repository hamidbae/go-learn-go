package response

import "net/http"

type InvalidArg struct {
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}

type Response struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Type       string      `json:"type"`
	Data       any         `json:"data,omitempty"`
	InvalidArg *InvalidArg `json:"invalid_arg,omitempty"`
}

type UsecaseError struct {
	HttpCode     int
	Message      string
	ErrorType    string
	Error error
}

var ResponseType = map[string]string{
	"00": "SUCCESS",
	"01": "ACCEPTED",
	"96": "BAD_REQUEST",
	"97": "UNAUTHENTICATED",
	"98": "FORBIDDEN",
	"99": "INTERNAL_SERVER_ERROR",
}

func Build(httpCode int, response Response) Response {
	switch httpCode {
	case http.StatusOK:
		response.Code = "00"
		response.Type = ResponseType[response.Code]
	case http.StatusAccepted:
		response.Code = "01"
		response.Type = ResponseType[response.Code]
	case http.StatusBadRequest:
		response.Code = "96"
		response.Type = ResponseType[response.Code]
	case http.StatusUnauthorized:
		response.Code = "97"
		response.Type = ResponseType[response.Code]
	case http.StatusForbidden:
		response.Code = "98"
		response.Type = ResponseType[response.Code]
	default:
		response.Code = "99"
		response.Type = ResponseType[response.Code]
	}

	return response
}