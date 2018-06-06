package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"
)

// Request defines the input type for demo HTTP inputs.
type Request struct {
	Start []string `json:"start"`
	End   []string `json:"end"`
}

// Response defines an output type for demo HTTP responses.
type Response struct {
	Combos []string `json:"combos" xml:"combo"`
}

// Controller is a simple demo implementation of a
// midl.MiddlewareFunc compatible controller which produces
// a cartesian product of input strings.
func Controller(req midl.Request) midl.Response {
	var input Request
	var output Response

	err := req.ProcessBody(midl.BodyProcessorFunc(func(in []byte) error {
		return json.Unmarshal(in, &input)
	})).Error()

	if err != nil {
		return midl.MakeErrorResponse(http.StatusInternalServerError, err)
	}

	for _, prefix := range input.Start {
		for _, suffix := range input.End {
			output.Combos = append(output.Combos, prefix+suffix)
		}
	}

	return midl.MakeResponse(http.StatusOK, output).
		SetHeader("Custom-Header", "some value")
}
