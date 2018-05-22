package main


import (
	"net/http"

	"github.com/foxcapades/go-midl/pkg/midl"
	gjs "github.com/xeipuuv/gojsonschema"
)

type Validator struct {
	schema *gjs.Schema
}

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
			out = append(out, v.Field() + ": " + v.Description())
		}
		return midl.MakeResponse(http.StatusBadRequest, out)
	}

	return nil
}
