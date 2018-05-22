package main

import (
	"log"
	"net/http"

	"github.com/foxcapades/go-midl/pkg/midl"
	"github.com/gorilla/mux"
	gjs "github.com/xeipuuv/gojsonschema"
)

const input = `{
  "type": "object",
  "properties": {
    "start": {
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "end":{
      "type": "array",
      "items": {
        "type": "string"
      }
    }
  },
  "required": [
    "start",
    "end"
  ]
}`

func main() {
	schema, err := gjs.NewSchema(gjs.NewStringLoader(input))
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	r.Handle("/combine", midl.XMLAdapter(
		Validator{schema},
		midl.MiddlewareFunc(Controller),
	)).Queries("xml", "").Methods(http.MethodPost)

	r.Handle("/combine", midl.JSONAdapter(
		Validator{schema},
		midl.MiddlewareFunc(Controller),
	)).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
