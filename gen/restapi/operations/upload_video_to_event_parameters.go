// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadVideoToEventParams creates a new UploadVideoToEventParams object
// no default values defined in spec.
func NewUploadVideoToEventParams() UploadVideoToEventParams {

	return UploadVideoToEventParams{}
}

// UploadVideoToEventParams contains all the bound params for the upload video to event operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadVideoToEvent
type UploadVideoToEventParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ID int64
	/*The file to upload.
	  Required: true
	  In: formData
	*/
	Upfile io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadVideoToEventParams() beforehand.
func (o *UploadVideoToEventParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	upfile, upfileHeader, err := r.FormFile("upfile")
	if err != nil {
		res = append(res, errors.New(400, "reading file %q failed: %v", "upfile", err))
	} else if err := o.bindUpfile(upfile, upfileHeader); err != nil {
		// Required: true
		res = append(res, err)
	} else {
		o.Upfile = &runtime.File{Data: upfile, Header: upfileHeader}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *UploadVideoToEventParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}

// bindUpfile binds file parameter Upfile.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UploadVideoToEventParams) bindUpfile(file multipart.File, header *multipart.FileHeader) error {
	return nil
}