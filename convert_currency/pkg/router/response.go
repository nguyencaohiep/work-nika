package router

import (
	"convert-service/pkg/log"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	Success             = "B.CODE.200"
	BadRequest          = "B.CODE.400"
	Unauthorized        = "B.CODE.401"
	TooManyRequests     = "B.CODE.429"
	NotFound            = "B.CODE.404"
	MethodNotAllowed    = "B.CODE.405"
	InternalServerError = "B.CODE.500"
)

type ResSuccess struct {
	Status  bool   `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResWithData struct {
	Status  bool   `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResError struct {
	Status  bool   `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func ResponseWrite(w http.ResponseWriter, responseCode int, responseData any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	json.NewEncoder(w).Encode(responseData)
}

func ResponseSuccess(w http.ResponseWriter, statusCode, message string) {
	var response ResSuccess

	if len(message) == 0 {
		message = "Success"
	}
	if len(statusCode) == 0 {
		statusCode = Success
	}

	response.Status = true
	response.Code = statusCode
	response.Message = message

	ResponseWrite(w, http.StatusOK, response)
}

func ResponseSuccessWithData(w http.ResponseWriter, statusCode, message string, data ...any) {
	var responseData any
	if len(data) == 1 {
		responseData = data[0]
	} else {
		responseData = data
	}
	var response ResWithData

	if len(message) == 0 {
		message = "Success"
	}

	if len(statusCode) == 0 {
		statusCode = Success
	}

	response.Status = true
	response.Code = statusCode
	response.Message = message
	response.Data = responseData

	ResponseWrite(w, http.StatusOK, response)
}

func ResponseCreatedWithData(w http.ResponseWriter, statusCode, message string, data ...any) {
	var responseData any
	if len(data) == 1 {
		responseData = data[0]
	} else {
		responseData = data
	}

	var response ResWithData
	if len(message) == 0 {
		message = "Success"
	}
	if len(statusCode) == 0 {
		statusCode = Success
	}

	response.Status = true
	response.Code = statusCode
	response.Message = message
	response.Data = responseData

	ResponseWrite(w, http.StatusCreated, response)
}

func ResponseBadRequest(w http.ResponseWriter, statusCode, message string) {
	var response ResError

	if len(message) == 0 {
		message = "Bad Request"
	}

	if len(statusCode) == 0 {
		statusCode = BadRequest
	}

	response.Status = false
	response.Code = statusCode
	response.Message = message
	response.Error = "Wrong request from client"

	ResponseWrite(w, http.StatusBadRequest, response)
}

func ResponseInternalError(w http.ResponseWriter, message string, err error) {
	var response ResError

	if len(message) == 0 {
		message = "Internal Server Error"
	}

	response.Status = false
	response.Code = InternalServerError
	response.Message = message
	response.Error = err.Error()

	log.Println(log.LogLevelError, message, strings.ToLower(err.Error()))

	ResponseWrite(w, http.StatusInternalServerError, response)
}

func ResponseUnauthorized(w http.ResponseWriter, message string) {
	var response ResError

	if len(message) == 0 {
		message = "Unauthorized"
	}

	response.Status = false
	response.Code = Unauthorized
	response.Message = "Unauthorized"
	response.Error = message

	ResponseWrite(w, http.StatusUnauthorized, response)
}

func ResponseTooManyRequests(w http.ResponseWriter) {
	var response ResError

	// Set Response Data
	response.Status = false
	response.Code = "B.ALL.429.C6"
	response.Message = "Too many requests"
	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusTooManyRequests, response)
}

func ResponseNotFound(w http.ResponseWriter, message string) {
	var response ResError

	if len(message) == 0 {
		message = "Not Found"
	}

	// Set Response Data
	response.Status = false
	response.Code = NotFound
	response.Message = "Not Found"
	response.Error = message

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusNotFound, response)
}

// ResponseMethodNotAllowed Function
func ResponseMethodNotAllowed(w http.ResponseWriter, message string) {
	var response ResError

	if len(message) == 0 {
		message = "Method Not Allowed"
	}

	// Set Response Data
	response.Status = false
	response.Code = MethodNotAllowed
	response.Message = "Method Not Allowed"
	response.Error = message

	// Set Response Data to HTTP
	ResponseWrite(w, http.StatusMethodNotAllowed, response)
}
