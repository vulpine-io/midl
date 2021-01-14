package midl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// NewAdapter creates a new Adapter instance with the provided settings
func NewStreamAdapter(
	content string,
	error ErrorSerializer,
	next ...Middleware,
) Adapter {
	return &streamAdapter{
		contentType:   content,
		errSerializer: error,
		emptyHandler:  DefaultEmptyHandler(),
		handlers:      next,
	}
}

type streamAdapter struct {
	handlers []Middleware
	wrappers []RequestWrapper

	errSerializer ErrorSerializer
	contentType   string
	emptyHandler  EmptyHandler
}

func (d streamAdapter) ServeHTTP(w writer, r *http.Request) {
	var res Response
	var wrapLen int

	req, err := NewRequest(r)
	if err != nil {
		d.writeError(w, err, req, NewResponse())
		return
	}

	wrapLen = len(d.wrappers)
	for i := 0; i < wrapLen; i++ {
		d.wrappers[i].Request(req)
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

	res = ensureResponse(res)

	for i := wrapLen - 1; i > -1; i-- {
		res = d.wrappers[i].Response(req, res)
	}

	res = ensureResponse(res)

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

func (d *streamAdapter) EmptyHandler(handler EmptyHandler) Adapter {
	d.emptyHandler = handler
	return d
}

func (d *streamAdapter) ContentType(contentType string) Adapter {
	d.contentType = contentType
	return d
}

func (d *streamAdapter) ErrorSerializer(err ErrorSerializer) Adapter {
	d.errSerializer = err
	return d
}

func (d *streamAdapter) Serializer(ser Serializer) Adapter {
	return d
}

func (d *streamAdapter) AddHandlers(mid ...Middleware) Adapter {
	d.handlers = append(d.handlers, mid...)
	return d
}

func (d *streamAdapter) SetHandlers(mid ...Middleware) Adapter {
	d.handlers = mid
	return d
}

func (d streamAdapter) writeEmpty(w writer, q Request, s Response) {
	body := d.emptyHandler.Handle(q, s)
	d.writeResponse(w, s.Code(), s.RawHeaders(), bytes.NewBuffer(body))
}

func (d streamAdapter) writeBody(w writer, _ Request, s Response) {
	var read io.Reader

	switch v := s.Body().(type) {
	case io.ReadCloser:
		defer v.Close()
		read = v
	case io.Reader:
		read = v
	case string:
		read = bytes.NewBufferString(v)
	case []byte:
		read = bytes.NewBuffer(v)
	default:
		read = bytes.NewBufferString(fmt.Sprint(v))
	}

	d.writeResponse(w, s.Code(), s.RawHeaders(), read)
}

func (d streamAdapter) writeError(w writer, e error, q Request, s Response) {
	body := d.errSerializer.Serialize(e, q, s)
	d.writeResponse(w, s.Code(), s.RawHeaders(), bytes.NewBuffer(body))
}

func (d streamAdapter) writeResponse(w writer, code int, head header, body io.Reader) {

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
		_, _ = io.Copy(w, body)
	}
}

func (d *streamAdapter) AddWrappers(w ...RequestWrapper) Adapter {
	d.wrappers = append(d.wrappers, w...)
	return d
}

func (d *streamAdapter) SetWrappers(w ...RequestWrapper) Adapter {
	d.wrappers = w
	return d
}
