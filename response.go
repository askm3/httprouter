package httprouter

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

const (
	badReqCode     = http.StatusBadRequest
	unauthCode     = http.StatusUnauthorized
	internalSECode = http.StatusInternalServerError
)

//Response a struct utility containing commons http responses
type ResponseHelper struct {
	headers map[string]string
}

//NewResponseHelper new instance with custom response headers
func NewResponseHelper(h map[string]string) *ResponseHelper {
	return &ResponseHelper{headers: h}
}

//BadRequest write a 400 header with body text "Bad Request"
func (rh *ResponseHelper) BadRequest(w http.ResponseWriter) error {
	return rh.StatusText(w, http.StatusText(badReqCode), badReqCode)
}

//OK write a 200 header with body encoded as json
func (rh *ResponseHelper)  OK(w http.ResponseWriter, res interface{}) error {
	return rh.Status(w, res, http.StatusOK)
}

//Unauthorized write a 401 header with body en
func (rh *ResponseHelper)  Unauthorized(w http.ResponseWriter) error {
	return rh.StatusText(w, http.StatusText(unauthCode), unauthCode)
}

//InternalServerError write a 500 header with body text "Internal Server Error"
func (rh *ResponseHelper) InternalServerError(w http.ResponseWriter) error {
	return rh.StatusText(w, http.StatusText(internalSECode), internalSECode)
}

//DbErr write a 404 header if there was a "No record found on db"
//or 500 header if there was a generic problem
func (rh *ResponseHelper)  DbErr(w http.ResponseWriter, err error) error {
	if err == sql.ErrNoRows {
		http.NotFound(w, nil)
		return nil
	}
	return rh.InternalServerError(w)
}

//Status to write custom header in the http response
func (rh *ResponseHelper)  StatusText(w http.ResponseWriter, res interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(&res)
}

//Status to write custom header in the http response
func (rh *ResponseHelper)  Status(w http.ResponseWriter, res interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(&res)
}
