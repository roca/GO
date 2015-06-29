// Package app provides application support for context and MongoDB access.
// Current Status Codes:
// 		200 OK           : StatusOK                  : Call is success and returning data.
//      204 No Content   : StatusNoContent           : Call is success and returns no data.
// 		400 Bad Request  : StatusBadRequest          : Invalid post data (syntax or semantics).
// 		401 Unauthorized : StatusUnauthorized        : Authentication failure.
// 		404 Not Found    : StatusNotFound            : Invalid URL or identifier.
// 		500 Internal     : StatusInternalServerError : Application specific beyond scope of user.
package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
)

// Invalid describes a validation error belonging to a specific field.
type Invalid struct {
	Fld string `json:"field_name"`
	Err string `json:"error"`
}

// jsonError is the response for errors that occur within the API.
type jsonError struct {
	Error  string    `json:"error"`
	Fields []Invalid `json:"fields,omitempty"`
}

// Context contains data associated with a single request.
type Context struct {
	Session *mgo.Session
	http.ResponseWriter
	Request   *http.Request
	Params    map[string]string
	SessionID string
}

// Error handles all error responses for the API.
func (c *Context) Error(err error) {
	switch err {
	case ErrNotFound:
		c.RespondError(err.Error(), http.StatusNotFound)
	case ErrInvalidID:
		c.RespondError(err.Error(), http.StatusBadRequest)
	case ErrValidation:
		c.RespondError(err.Error(), http.StatusBadRequest)
	default:
		c.RespondError(err.Error(), http.StatusInternalServerError)
	}
}

// Respond sends JSON to the client.
// If code is StatusNoContent, v is expected to be nil.
func (c *Context) Respond(v interface{}, code int) {
	log.Printf("%v : api : Respond [%d] : Started", c.SessionID, code)

	if code == http.StatusNoContent {
		c.WriteHeader(http.StatusNoContent)
		return
	}

	data, err := json.Marshal(v)
	if err != nil {
		// We want this error condition to panic so we get a stack trace. This should
		// never happen. The http package will catch the panic and provide logging
		// and return a 500 back to the caller.
		log.Panicf("%v : api : Respond [%d] : Failed: %v", c.SessionID, code, err)
	}

	datalen := len(data) + 1 // account for trailing LF
	h := c.Header()
	h.Set("Content-Type", "application/json")
	h.Set("Content-Length", strconv.Itoa(datalen))
	c.WriteHeader(code)
	fmt.Fprintf(c, "%s\n", data)

	log.Printf("%v : api : Respond [%d] : Completed", c.SessionID, code)
}

// RespondInvalid sends JSON describing field validation errors.
func (c *Context) RespondInvalid(fields []Invalid) {
	v := jsonError{
		Error:  "field validation failure",
		Fields: fields,
	}
	c.Respond(v, http.StatusBadRequest)
}

// RespondError sends JSON describing the error
func (c *Context) RespondError(error string, code int) {
	c.Respond(jsonError{Error: error}, code)
}
