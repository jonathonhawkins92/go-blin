package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/runtime/middleware"

	"goblin/client/responses"
)

type Client struct {
	mux      *http.ServeMux
	openapi  openapi3.T
	routes   []Route
	handlers map[string]map[string]Handler
}

type Route struct {
	Method    string
	Path      string
	Handler   Handler
	Operation *openapi3.Operation
}

func API(info openapi3.Info) *Client {
	return &Client{
		mux: http.NewServeMux(),
		openapi: openapi3.T{
			OpenAPI: "3.0.0",
			Info:    &info,
			Paths:   &openapi3.Paths{},
		},
		routes:   make([]Route, 0),
		handlers: make(map[string]map[string]Handler, 0),
	}
}

func (api *Client) createHandler(route Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler, ok := api.handlers[route.Path][r.Method]
		if !ok {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

		value := reflect.ValueOf(handler)
		results := value.Call([]reflect.Value{reflect.ValueOf(r)})

		if len(results) != 2 {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response, ok := results[0].Interface().(responses.IResponse)
		if !ok {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err, ok := results[1].Interface().(error)
		if ok && err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", response.ContentType())
		w.WriteHeader(response.Status())

		if response.Body() == nil {
			return
		}
		if err := response.Encode(w); err != nil {
			// Log the error, but we can't change the response at this point
			// as headers and status code have already been sent
			fmt.Printf("Error encoding response: %v\n", err)
		}
	}
}

func (api *Client) RegisterHandler(route Route) {
	if api.handlers[route.Path] == nil {
		api.handlers[route.Path] = make(map[string]Handler, 0)
	}
	api.handlers[route.Path][route.Method] = route.Handler
}

// Serve starts the HTTP server for the API
func (api *Client) Serve(port int) error {
	err := api.openapi.Validate(context.TODO())
	if err != nil {
		log.Fatalf("invalid OpenAPI spec: %v", err)
	}

	for _, route := range api.routes {
		api.mux.HandleFunc(route.Path, api.createHandler(route))
	}

	// Add OpenAPI JSON endpoint
	api.mux.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(api.openapi)
	})

	// Setup Swagger UI
	opts := middleware.SwaggerUIOpts{SpecURL: "/openapi.json"}
	sh := middleware.SwaggerUI(opts, nil)
	api.mux.Handle("/docs", sh)

	// Helpful logs
	print("---\n")
	log.Printf("Server running on http://localhost:%d", port)
	log.Printf("Swagger UI available at http://localhost:%d/docs", port)

	// Start server
	_port := strconv.Itoa(port)
	return http.ListenAndServe(":"+_port, api.mux)
}

func (api *Client) Endpoint(method string, path string, handler Handler) {
	route := Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
	api.RegisterHandler(route)
	api.routes = append(api.routes, route)
}

// Get adds a GET route to the Client
func (api *Client) Get(path string, handler Handler) {
	api.Endpoint(http.MethodGet, path, handler)
}

// Head adds a HEAD route to the Client
func (api *Client) Head(path string, handler Handler) {
	api.Endpoint(http.MethodHead, path, handler)
}

// Post adds a POST route to the Client
func (api *Client) Post(path string, handler Handler) {
	api.Endpoint(http.MethodPost, path, handler)
}

// Put adds a PUT route to the Client
func (api *Client) Put(path string, handler Handler) {
	api.Endpoint(http.MethodPut, path, handler)
}

// Patch adds a PATCH route to the Client
func (api *Client) Patch(path string, handler Handler) {
	api.Endpoint(http.MethodPatch, path, handler)
}

// Delete adds a DELETE route to the Client
func (api *Client) Delete(path string, handler Handler) {
	api.Endpoint(http.MethodDelete, path, handler)
}

// Connect adds a CONNECT route to the Client
func (api *Client) Connect(path string, handler Handler) {
	api.Endpoint(http.MethodConnect, path, handler)
}

// Options adds an OPTIONS route to the Client
func (api *Client) Options(path string, handler Handler) {
	api.Endpoint(http.MethodOptions, path, handler)
}

// Trace adds a TRACE route to the Client
func (api *Client) Trace(path string, handler Handler) {
	api.Endpoint(http.MethodTrace, path, handler)
}

type Handler func(r *http.Request) (responses.IResponse, error)

type ConfigResponse struct {
	Description *string             `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]string   `json:"headers,omitempty" yaml:"headers,omitempty"`
	Model       responses.IResponse `json:"model,omitempty" yaml:"model,omitempty"`
	Links       map[string]string   `json:"links,omitempty" yaml:"links,omitempty"`
	Validator   func(value interface{}) bool
}

type Config struct {
	Path        string
	Method      string
	Summary     string
	Description string
	Tags        []string
	Responses   map[int]ConfigResponse
}

func ConfigToOpenapi3(config Config) *openapi3.Operation {
	base := openapi3.Operation{}
	if config.Tags != nil && len(config.Tags) > 0 {
		base.Tags = config.Tags
	}
	if config.Description != "" {
		base.Description = config.Description
	}
	return &base
}

func (route *Route) Summary(value string) *Route {
	route.Operation.Summary = value
	return route
}
func (route *Route) Description(value string) *Route {
	route.Operation.Description = value
	return route
}

// For custom responses, the user will need to do blin.RegisterResponse(response IRepsonse)
// The Resposes will be used when generating swagger docs
// If not registered everything will work but the respons wont be in the docs
