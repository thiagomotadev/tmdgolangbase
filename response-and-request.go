package tmdgolangbase

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// RequestAndResponse ..
type RequestAndResponse struct {
	w http.ResponseWriter
	r *http.Request
}

func (rar RequestAndResponse) encode(data interface{}) (err error) {
	err = json.NewEncoder(rar.w).Encode(data)
	return
}

func (rar RequestAndResponse) decode(data interface{}) (err error) {
	err = json.NewDecoder(rar.r.Body).Decode(data)
	return
}

// GetRequestData ...
func (rar RequestAndResponse) GetRequestData(data interface{}) (err error) {
	err = rar.decode(data)
	return
}

// SetResponseStatus ...
func (rar RequestAndResponse) SetResponseStatus(status int) {
	rar.w.WriteHeader(status)
}

// SetResponseData ...
func (rar RequestAndResponse) SetResponseData(status int, data interface{}) (err error) {
	rar.SetResponseStatus(status)
	rar.w.Header().Add("Content-Type", "application/json")
	err = rar.encode(data)
	return
}

func (rar RequestAndResponse) setResponseError(err error) {
	type responseData struct {
		Message string
	}

	response := responseData{
		Message: err.Error(),
	}

	rar.SetResponseStatus(http.StatusBadRequest)
	rar.w.Header().Add("Content-Type", "application/json")
	rar.encode(response)
}

// GetRequestParameter ...
func (rar RequestAndResponse) GetRequestParameter(name string) (value string, err error) {
	vars := mux.Vars(rar.r)
	value = vars[name]
	return
}

// GetRequestUintParameter ...
func (rar RequestAndResponse) GetRequestUintParameter(name string) (value uint, err error) {
	vars := mux.Vars(rar.r)
	valueString := vars[name]
	temp, err := strconv.ParseUint(valueString, 10, 32)
	value = uint(temp)
	return
}
