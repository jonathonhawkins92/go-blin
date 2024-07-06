package responses

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
)

type IResponse interface {
	Body() any
	ContentType() string
	Status() int
	Encode(w io.Writer) error
}

type Encoder func(w io.Writer, v any) error

type Response struct {
	body        any
	status      int
	contentType string
	encoder     Encoder
}

func (response *Response) Body() any {
	return response.body
}
func (response *Response) ContentType() string {
	return response.contentType
}
func (response *Response) Status() int {
	return response.status
}

func (response *Response) Encode(w io.Writer) error {
	return response.encoder(w, response.body)
}

type EncoderStrategy struct {
	ContentType string
	Encode      Encoder
}

// RegisterEncoder registers a new encoder strategy
func (response *Response) RegisterEncoder(strategy EncoderStrategy) {
	// Consider checking for duplicates or overwrites here if needed
	response.encoder = strategy.Encode
	response.contentType = strategy.ContentType
}

// JSON sets the content type to JSON and uses json.NewEncoder(w).Encode as the encoder
func (response *Response) JSON(body interface{}) *Response {
	response.body = body
	response.RegisterEncoder(EncoderStrategy{
		ContentType: "application/json",
		Encode: func(w io.Writer, v any) error {
			return json.NewEncoder(w).Encode(v)
		},
	})
	return response
}

// XML sets the content type to XML and uses xml.NewEncoder(w).Encode as the encoder
func (response *Response) XML(body interface{}) *Response {
	response.body = body
	response.RegisterEncoder(EncoderStrategy{
		ContentType: "application/xml",
		Encode: func(w io.Writer, v any) error {
			return xml.NewEncoder(w).Encode(v)
		},
	})
	return response
}

// Text sets the content type to plain text and uses a simple string conversion as the encoder
func (response *Response) Text(body string) *Response {
	response.body = body
	response.RegisterEncoder(EncoderStrategy{
		ContentType: "text/plain",
		Encode: func(w io.Writer, v any) error {
			_, err := fmt.Fprint(w, v)
			return err
		},
	})
	return response
}
