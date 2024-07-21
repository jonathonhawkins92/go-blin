package goblin

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/runtime/middleware"

	"goblin/responses"
)

func init() {
	log.Println("BUILDING?")
	// // Set up the custom import hook
	// originalImportHook := build.Default.ImportHook
	// build.Default.ImportHook = func(dir string, importPath string, fromDir string, mode build.ImportMode) (*build.Package, error) {
	// 	// Call the original ImportHook
	// 	pkg, err := originalImportHook(dir, importPath, fromDir, mode)
	// 	if err != nil {
	// 		return pkg, err
	// 	}

	// 	// Perform our custom analysis
	// 	analyzePackage(pkg)

	// 	return pkg, nil
	// }
}

type Client struct {
	mux       *http.ServeMux
	openapi   openapi3.T
	routes    []Route
	handlers  map[string]map[string]Handler
	responses map[string]map[string]responses.Response
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
		routes:    make([]Route, 0),
		handlers:  make(map[string]map[string]Handler, 0),
		responses: make(map[string]map[string]responses.Response, 0),
	}
}

func (api *Client) createHandler(route Route, responseTemplate *responses.Response) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		handler, ok := api.handlers[route.Path][request.Method]
		if !ok {
			http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
		}
		value := reflect.ValueOf(handler)
		results := value.Call([]reflect.Value{
			reflect.ValueOf(request),
			reflect.ValueOf(responseTemplate),
		})

		if len(results) != 2 {
			http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response, ok := results[0].Interface().(responses.IResponse)
		if !ok {
			http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err, ok := results[1].Interface().(error)
		if ok && err != nil {
			log.Println(err.Error())
			http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		responseWriter.Header().Set("Content-Type", response.ContentType())
		responseWriter.WriteHeader(response.Status())

		if response.Body() == nil {
			return
		}
		if err := response.Encode(responseWriter); err != nil {
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
		// Idea: by making the response here and passing it forward we know what the
		// memory address the response is so can use it as a marker when analyzing the
		// users handler code to find the responses from it.
		response := responses.NewResponse()
		// valueOf := reflect.ValueOf(response)
		// typeOf := valueOf.Type()
		// log.Println("response", valueOf, typeOf, valueOf.Pointer())
		// responseGenerator.AnalyzeFunction(route.Handler, response)
		api.mux.HandleFunc(route.Path, api.createHandler(route, response))
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

func (api *Client) Endpoint(method string, config Config, handler Handler) {
	route := Route{
		Method:  method,
		Path:    config.Path,
		Handler: handler,
	}
	if pathItem := api.openapi.Paths.Value(config.Path); pathItem == nil {
		newPathItem := openapi3.PathItem{}
		api.openapi.Paths.Set(config.Path, &newPathItem)
	}
	pathItem := api.openapi.Paths.Value(config.Path)
	if currrentOperation := pathItem.GetOperation(method); currrentOperation != nil {
		slog.Warn("Method already exsists on route, overwriting.")
	}
	operation := ConfigToOpenapi3(config)
	pathItem.SetOperation(method, operation)
	api.RegisterHandler(route)
	api.routes = append(api.routes, route)
}

// Get adds a GET route to the Client
func (api *Client) Get(config Config, handler Handler) {
	api.Endpoint(http.MethodGet, config, handler)
}

// Head adds a HEAD route to the Client
func (api *Client) Head(config Config, handler Handler) {
	api.Endpoint(http.MethodHead, config, handler)
}

// Post adds a POST route to the Client
func (api *Client) Post(config Config, handler Handler) {
	api.Endpoint(http.MethodPost, config, handler)
}

// Put adds a PUT route to the Client
func (api *Client) Put(config Config, handler Handler) {
	api.Endpoint(http.MethodPut, config, handler)
}

// Patch adds a PATCH route to the Client
func (api *Client) Patch(config Config, handler Handler) {
	api.Endpoint(http.MethodPatch, config, handler)
}

// Delete adds a DELETE route to the Client
func (api *Client) Delete(config Config, handler Handler) {
	api.Endpoint(http.MethodDelete, config, handler)
}

// Connect adds a CONNECT route to the Client
func (api *Client) Connect(config Config, handler Handler) {
	api.Endpoint(http.MethodConnect, config, handler)
}

// Options adds an OPTIONS route to the Client
func (api *Client) Options(config Config, handler Handler) {
	api.Endpoint(http.MethodOptions, config, handler)
}

// Trace adds a TRACE route to the Client
func (api *Client) Trace(config Config, handler Handler) {
	api.Endpoint(http.MethodTrace, config, handler)
}

type Handler func(request *http.Request, response *responses.Response) (*responses.Response, error)

type ConfigResponse struct {
	Description *string           `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	Model       any               `json:"model,omitempty" yaml:"model,omitempty"`
	Links       map[string]string `json:"links,omitempty" yaml:"links,omitempty"`
	Validator   func(value any) bool
}

type Config struct {
	Path        string
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
	base.Responses = openapi3.NewResponses(
		openapi3.WithStatus(200, &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("OK").
				WithContent(
					openapi3.NewContentWithSchema(
						openapi3.NewStringSchema(),
						[]string{"text/plain"},
					),
				),
		}),
	)
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
