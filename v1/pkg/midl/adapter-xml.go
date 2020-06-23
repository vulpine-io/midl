package midl

import "encoding/xml"

// XMLAdapter creates a new Adapter defaulted for XML responses
func XMLAdapter(handlers ...Middleware) Adapter {
	return &adapter{
		contentType:   "application/xml",
		serializer:    SerializerFunc(xml.Marshal),
		errSerializer: DefaultXMLErrorSerializer(),
		emptyHandler:  DefaultEmptyHandler(),
		handlers:      handlers,
	}
}
