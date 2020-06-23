package midl

import "encoding/json"

// JSONAdapter creates a new Adapter defaulted for JSON responses
func JSONAdapter(handlers ...Middleware) Adapter {
	return &adapter{
		contentType:   "application/json",
		serializer:    SerializerFunc(json.Marshal),
		errSerializer: DefaultJSONErrorSerializer(),
		emptyHandler:  DefaultEmptyHandler(),
		handlers:      handlers,
	}
}
