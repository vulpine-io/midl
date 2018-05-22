package main

import (
	"encoding/json"
	"net/http"

	"github.com/foxcapades/go-midl/pkg/midl"
)

type Request struct {
	Start []string `json:"start"`
	End   []string `json:"end"`
}

type Response struct {
	Combos []string `json:"combos" xml:"combo"`
}

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
