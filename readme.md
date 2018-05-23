# midl



![license](https://img.shields.io/github/license/Foxcapades/go-midl.svg)
[![Build Status](https://travis-ci.org/Foxcapades/go-midl.svg?branch=master)](https://travis-ci.org/Foxcapades/go-midl)
[![Codecov](https://img.shields.io/codecov/c/github/Foxcapades/go-midl.svg)](https://codecov.io/gh/Foxcapades/go-midl)
[![Go Report Card](https://goreportcard.com/badge/github.com/Foxcapades/go-midl)](https://goreportcard.com/report/github.com/Foxcapades/go-midl)


Serializing middleware layer

```bash
$ go get github.com/foxcapades/go-midl/pkg/midl
```

## Example

A full working example can be found in the cmd/test package.

```go
func main() {
	http.Handle("/", midl.JSONAdapter(
		SessionValidator(),
		InputValidator(),
		Controller(),
	))
	log.Fatal(http.ListenAndServe(":80", nil))
}
``` 
