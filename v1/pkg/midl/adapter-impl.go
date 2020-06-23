package midl

import "net/http"

// NewAdapter creates a new Adapter instance with the provided settings
func NewAdapter(
	content string,
	serial Serializer,
	error ErrorSerializer,
	next ...Middleware,
) Adapter {
	return &adapter{
		contentType:   content,
		serializer:    serial,
		errSerializer: error,
		emptyHandler:  DefaultEmptyHandler(),
		handlers:      next,
	}
}

type adapter struct {
	handlers      []Middleware
	serializer    Serializer
	errSerializer ErrorSerializer
	contentType   string
	emptyHandler  EmptyHandler
}

func (d adapter) ServeHTTP(w writer, r *http.Request) {
	var res Response

	req, err := NewRequest(r)
	if err != nil {
		d.writeError(w, err, req, NewResponse())
		return
	}

	for _, hand := range d.handlers {
		if hand == nil {
			continue
		}

		res = hand.Handle(req)

		if res != nil {
			break
		}
	}

	if res == nil {
		d.writeError(w, ErrNoHandlers, req, NewResponse())
		return
	}

	if res.Error() != nil {
		d.writeError(w, res.Error(), req, res)
		return
	}

	if res.Body() == nil {
		d.writeEmpty(w, req, res)
		return
	}

	d.writeBody(w, req, res)

	for _, fn := range res.Callbacks() {
		go fn()
	}
}

func (d *adapter) EmptyHandler(handler EmptyHandler) Adapter {
	d.emptyHandler = handler
	return d
}

func (d *adapter) ContentType(contentType string) Adapter {
	d.contentType = contentType
	return d
}

func (d *adapter) ErrorSerializer(err ErrorSerializer) Adapter {
	d.errSerializer = err
	return d
}

func (d *adapter) Serializer(ser Serializer) Adapter {
	d.serializer = ser
	return d
}

func (d *adapter) AddHandlers(mid ...Middleware) Adapter {
	d.handlers = append(d.handlers, mid...)
	return d
}

func (d *adapter) SetHandlers(mid ...Middleware) Adapter {
	d.handlers = mid
	return d
}

func (d adapter) writeEmpty(w writer, q Request, s Response) {
	body := d.emptyHandler.Handle(q, s)
	d.writeResponse(w, s.Code(), s.RawHeaders(), body)
}

func (d adapter) writeBody(w writer, q Request, s Response) {
	body, err := d.serializer.Serialize(s.Body())

	if err != nil {
		d.writeError(w, err, q, s)
		return
	}

	d.writeResponse(w, s.Code(), s.RawHeaders(), body)
}

func (d adapter) writeError(w writer, e error, q Request, s Response) {
	body := d.errSerializer.Serialize(e, q, s)
	d.writeResponse(w, s.Code(), s.RawHeaders(), body)
}

func (d adapter) writeResponse(w writer, code int, head header, body []byte) {

	// Don't override user provided header if present.
	if _, ok := head["Content-Type"]; !ok && d.contentType != "" {
		w.Header().Set("Content-Type", d.contentType)
	}

	if head != nil {
		for key, values := range head {
			for _, val := range values {
				w.Header().Add(key, val)
			}
		}
	}

	w.WriteHeader(code)
	if body == nil {
		_, _ = w.Write([]byte{})
	} else {
		_, _ = w.Write(body)
	}
}
