# midl

Serializing middleware layer

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
