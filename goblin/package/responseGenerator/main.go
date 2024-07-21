package main

import (
	"encoding/json"
	"fmt"
	"goblin/responses"
	"io"
	"net/http"
	"sync"
)

func AnalyzeFunction(fn interface{}, response *responses.Response) {}

// I've got an struct called Response, I pass it into a function as a part of creating a endpoint along with the function that takes, the Response struct is used by my user to create a valid response object including set things like the content type, because this can happen in multiple paths with in the users code the only way I think I can find how they use it fully is to create an AST of the code they pass me and find references to the Response object, is that accurate?

type ResponseBuilder struct {
	statusCode int
	headers    map[string]string
	body       interface{}
	operations []func(*ResponseBuilder)
	mu         sync.Mutex
	usageLog   []string
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		headers:    make(map[string]string),
		operations: make([]func(*ResponseBuilder), 0),
		usageLog:   make([]string, 0),
	}
}

func (rb *ResponseBuilder) logUsage(method string) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	rb.usageLog = append(rb.usageLog, method)
}

func (rb *ResponseBuilder) HTTP200() *ResponseBuilder {
	rb.logUsage("HTTP200")
	rb.operations = append(rb.operations, func(r *ResponseBuilder) {
		r.statusCode = http.StatusOK
	})
	return rb
}
func (rb *ResponseBuilder) HTTP404() *ResponseBuilder {
	rb.logUsage("HTTP404")
	rb.operations = append(rb.operations, func(r *ResponseBuilder) {
		r.statusCode = http.StatusNotFound
	})
	return rb
}

func (rb *ResponseBuilder) JSON(data interface{}) *ResponseBuilder {
	rb.logUsage("JSON")
	rb.operations = append(rb.operations, func(r *ResponseBuilder) {
		r.headers["Content-Type"] = "application/json"
		r.body = data
	})
	return rb
}
func (rb *ResponseBuilder) PlainText(data interface{}) *ResponseBuilder {
	rb.logUsage("PlainText")
	rb.operations = append(rb.operations, func(r *ResponseBuilder) {
		r.headers["Content-Type"] = "text/plain"
		r.body = data
	})
	return rb
}

// Add more methods as needed...

func (rb *ResponseBuilder) Build() *http.Response {
	for _, op := range rb.operations {
		op(rb)
	}

	// Create the actual response
	response := &http.Response{
		StatusCode: rb.statusCode,
		Header:     make(http.Header),
	}

	for k, v := range rb.headers {
		response.Header.Set(k, v)
	}

	if rb.body != nil {
		bodyBytes, _ := json.Marshal(rb.body)
		response.Body = &responseBody{body: bodyBytes}
	}

	return response
}

func (rb *ResponseBuilder) GetUsageLog() []string {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	return append([]string{}, rb.usageLog...)
}

// Helper struct to implement io.ReadCloser for response body
type responseBody struct {
	body []byte
	idx  int64
}

func (rb *responseBody) Read(p []byte) (n int, err error) {
	if rb.idx >= int64(len(rb.body)) {
		return 0, io.EOF
	}
	n = copy(p, rb.body[rb.idx:])
	rb.idx += int64(n)
	return
}

func (rb *responseBody) Close() error {
	return nil
}

type User struct {
	ID   int
	Name string
}

// Example usage
func main() {
	builder := NewResponseBuilder()

	if true == true {
		builder.HTTP200().JSON(User{ID: 1, Name: "John"})
	} else {
		builder.HTTP404().PlainText("User not found")
	}

	response := builder.Build()

	fmt.Println("Usage log:", builder.GetUsageLog())
	fmt.Println(response)
}
