// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WebOKCode is the HTTP code returned for type WebOK
const WebOKCode int = 200

/*
WebOK Page found

swagger:response webOK
*/
type WebOK struct {
}

// NewWebOK creates WebOK with default headers values
func NewWebOK() *WebOK {

	return &WebOK{}
}

// WriteResponse to the client
func (o *WebOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WebNotFoundCode is the HTTP code returned for type WebNotFound
const WebNotFoundCode int = 404

/*
WebNotFound Resource not found.

swagger:response webNotFound
*/
type WebNotFound struct {
}

// NewWebNotFound creates WebNotFound with default headers values
func NewWebNotFound() *WebNotFound {

	return &WebNotFound{}
}

// WriteResponse to the client
func (o *WebNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}