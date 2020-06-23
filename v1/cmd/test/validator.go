package main

import (
	"net/http"

	"github.com/vulpine-io/midl/v1/pkg/midl"
	gjs "github.com/xeipuuv/gojsonschema"
)

// Validator is a midl.Middleware implementation that
// validates incoming request bodies against a JSON schema
// provided with its construction.
type Validator struct {
	schema *gjs.Schema
}

// Handle will check the request body against the JSON
// schema stored on the Validator struct and return a
// response if an error occurs or if the request is invalid.
// If the request is valid it will return nothing, which
// will signal to the Adapter that the request has not yet
// been handled.
func (v Validator) Handle(req midl.Request) midl.Response {
	var res *gjs.Result

	req.ProcessBody(midl.BodyProcessorFunc(func(in []byte) (err error) {
		res, err = v.schema.Validate(gjs.NewBytesLoader(req.Body()))
		return
	}))

	if req.Error() != nil {
		return midl.MakeErrorResponse(http.StatusInternalServerError, req.Error())
	}

	if !res.Valid() {
		var out []string
		for _, v := range res.Errors() {
			out = append(out, v.Field()+": "+v.Description())
		}
		return midl.MakeResponse(http.StatusBadRequest, out)
	}

	return nil
}
