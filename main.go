package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-openapi/runtime/middleware"
)

// User represents a user in the system
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Operation   *openapi3.Operation
}

// Define your routes here
var routes = []Route{
	{
		Method:  http.MethodGet,
		Pattern: "/hello",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		},
		Operation: &openapi3.Operation{
			Summary:     "Say Hello",
			Description: "Returns a hello message",
			Responses: openapi3.NewResponses(
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
			),
		},
	},
	{
		Method:  http.MethodGet,
		Pattern: "/users/{id}",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			parts := strings.Split(r.URL.Path, "/")
			if len(parts) != 3 {
				http.Error(w, "Invalid URL", http.StatusBadRequest)
				return
			}
			userID := parts[2]
			if userID == "1" {
				user := User{ID: 1, Name: "John Doe", Age: 30}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(user)
			} else {
				http.Error(w, "User not found", http.StatusNotFound)
			}
		},
		Operation: &openapi3.Operation{
			Summary:     "Get User",
			Description: "Returns a user by ID",
			Parameters: openapi3.Parameters{
				&openapi3.ParameterRef{
					Value: openapi3.NewPathParameter("id").
						WithSchema(openapi3.NewIntegerSchema()),
				},
			},
			Responses: openapi3.NewResponses(
				openapi3.WithStatus(200, &openapi3.ResponseRef{
					Value: openapi3.NewResponse().
						WithDescription("OK").
						WithContent(
							openapi3.NewContentWithSchemaRef(
								openapi3.NewSchemaRef("#/components/schemas/User", &openapi3.Schema{}),
								[]string{"application/json"},
							),
						),
				}),
				openapi3.WithStatus(404, &openapi3.ResponseRef{
					Value: openapi3.NewResponse().
						WithDescription("Not Found"),
				}),
			),
		},
	},
	{
		Method:  http.MethodPost,
		Pattern: "/users",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			user.ID = 1 // Example assignment of user ID
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(user)
		},
		Operation: &openapi3.Operation{
			Summary:     "Create User",
			Description: "Creates a new user",
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Required: true,
					Content: openapi3.NewContentWithSchemaRef(
						openapi3.NewSchemaRef("#/components/schemas/User", &openapi3.Schema{}),
						[]string{"application/json"},
					),
				},
			},
			Responses: openapi3.NewResponses(
				openapi3.WithStatus(201, &openapi3.ResponseRef{
					Value: openapi3.NewResponse().
						WithDescription("User created").
						WithContent(
							openapi3.NewContentWithSchemaRef(
								openapi3.NewSchemaRef("#/components/schemas/User", &openapi3.Schema{}),
								[]string{"application/json"},
							),
						),
				}),
				openapi3.WithStatus(400, &openapi3.ResponseRef{
					Value: openapi3.NewResponse().
						WithDescription("Invalid input"),
				}),
			),
		},
	},
}

func generateSchema(v interface{}) *openapi3.Schema {
	t := reflect.TypeOf(v)
	schema := &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: make(openapi3.Schemas),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}

		var fieldSchema *openapi3.Schema
		switch field.Type.Kind() {
		case reflect.Int:
			fieldSchema = openapi3.NewIntegerSchema()
		case reflect.String:
			fieldSchema = openapi3.NewStringSchema()
		// Add more cases here for other types as needed
		default:
			// For unsupported types, you might want to log a warning or handle differently
			fieldSchema = openapi3.NewSchema()
		}

		schema.Properties[jsonTag] = openapi3.NewSchemaRef("", fieldSchema)
	}

	return schema
}

// setupOpenAPI initializes the OpenAPI documentation
func setupOpenAPI() *openapi3.T {
	openapi := openapi3.T{}
	openapi.OpenAPI = "3.0.0"
	openapi.Info = &openapi3.Info{
		Title:   "Example API",
		Version: "0.0.1",
	}
	openapi.Servers = openapi3.Servers{
		&openapi3.Server{
			URL: "http://localhost:8080",
		},
	}

	// Generate schema for User
	userSchema := generateSchema(User{})

	openapi.Components = &openapi3.Components{
		Schemas: openapi3.Schemas{
			"User": openapi3.NewSchemaRef("", userSchema),
		},
	}

	openapi.Paths = openapi3.NewPaths()
	for _, route := range routes {
		pathItem := openapi3.PathItem{}
		switch route.Method {
		case http.MethodGet:
			pathItem.Get = route.Operation
		case http.MethodPost:
			pathItem.Post = route.Operation
		case http.MethodPut:
			pathItem.Put = route.Operation
		case http.MethodDelete:
			pathItem.Delete = route.Operation
		}

		openapi.Paths.Set(route.Pattern, &pathItem)
	}

	return &openapi
}

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		remoteAddr := r.RemoteAddr
		if ip := r.Header.Get("X-Real-IP"); ip != "" {
			remoteAddr = ip
		} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
			remoteAddr = strings.Split(ip, ", ")[0]
		}

		log.Printf(
			`"%s %s %s" from %s - %d %dB in %v`,
			r.Method,
			r.URL.RequestURI(),
			r.Proto,
			remoteAddr,
			rw.status,
			rw.size,
			duration,
		)
	})
}

// matchRoute checks if the request URL matches the route pattern
func matchRoute(path string, pattern string) bool {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	urlParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(patternParts) != len(urlParts) {
		return false
	}

	for i, part := range patternParts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			// This is a path parameter, it matches any value
			continue
		}
		if part != urlParts[i] {
			return false
		}
	}

	return true
}

func main() {
	// OpenAPI setup
	openapi := setupOpenAPI()

	// Validate OpenAPI spec
	err := openapi.Validate(context.TODO())
	if err != nil {
		log.Fatalf("invalid OpenAPI spec: %v", err)
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Add OpenAPI JSON endpoint
	mux.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(openapi)
	})

	// Setup Swagger UI
	opts := middleware.SwaggerUIOpts{SpecURL: "/openapi.json"}
	sh := middleware.SwaggerUI(opts, nil)
	mux.Handle("/docs", sh)

	// Register a single handler for all routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("THIS", r.URL.Path)
		for _, route := range routes {
			if matchRoute(r.URL.Path, route.Pattern) && r.Method == route.Method {
				route.HandlerFunc(w, r)
				return
			}
		}
		http.NotFound(w, r)
	})

	// Wrap the ServeMux with the logging middleware
	handler := loggingMiddleware(mux)

	// Start the server
	port := 8080
	log.Printf("Server running on http://localhost:%d", port)
	log.Printf("Swagger UI available at http://localhost:%d/docs", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
