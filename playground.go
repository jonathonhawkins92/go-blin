package main

import (
	"encoding/json"
	"fmt"
	goblin "goblin/client"
	"goblin/client/responses"
	"io"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// PokemonResponse represents the structure of the Pokemon API response
type PokemonResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

// fetchPokemonData retrieves data for a given Pokemon from the PokeAPI
func fetchPokemonData(pokemonName string) (*PokemonResponse, error) {
	// Make an HTTP GET request to the Pokemon API
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Create a PokemonResponse struct to hold the unmarshaled data
	var pokemon PokemonResponse

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &pokemon, nil
}
func main() {

	api := goblin.API(openapi3.Info{
		Title:       "My example API",
		Version:     "0.0.1",
		Description: "My cool app",
	})

	api.Get("/", func(r *http.Request) (responses.IResponse, error) {
		return responses.Http200().Text("hi"), nil
	})
	api.Get("/example", func(r *http.Request) (responses.IResponse, error) {
		pokemon, err := fetchPokemonData("pikachu")
		if err != nil {
			return responses.Http500().Text("bang"), nil
		}
		return responses.Http200().JSON(pokemon), nil
	})

	api.Put("/my-orbies", func(req *http.Request) (responses.IResponse, error) {
		return responses.Continue().Text("hi"), nil
	})

	// api.Get("/example2", func(r *http.Request) (responses.IResponse, error) {
	// 	if x == "bad" {
	// 		return responses.BadRequest("lol really?"), nil
	// 	}
	// 	return &responses.Response{}, nil
	// })

	// api.Put(Config{
	// 	Path:        "/my/path/{id}",
	// 	Description: "Send me some really nice data!",
	// }, func(request *http.Request) (responses.IResponse, error) {
	// 	if x == "bad" {
	// 		return responses.BadRequest("lol really?"), nil
	// 	} else if x == "err" {
	// 		return responses.InternalServerError("bang"), nil
	// 	} else {
	// 		pokemon, err := fetchPokemonData("pikachu")
	// 		if err != nil {
	// 			return responses.InternalServerError("bang"), nil
	// 		}
	// 		return responses.OK(pokemon).JSON(), nil
	// 	}
	// })

	// Validate OpenAPI spec
	// err := api.openapi.Validate(context.TODO())
	// if err != nil {
	// 	log.Fatalf("invalid OpenAPI spec: %v", err)
	// }

	// Create a new ServeMux
	// mux := http.NewServeMux()

	// // Add OpenAPI JSON endpoint
	// mux.HandleFunc("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(openapi)
	// })

	// // Setup Swagger UI
	// opts := middleware.SwaggerUIOpts{SpecURL: "/openapi.json"}
	// sh := middleware.SwaggerUI(opts, nil)
	// mux.Handle("/docs", sh)

	// Register a single handler for all routes
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("THIS", r.URL.Path)
	// 	for _, route := range routes {
	// 		if matchRoute(r.URL.Path, route.Path) && r.Method == route.Method {
	// 			route.HandlerFunc(w, r)
	// 			return
	// 		}
	// 	}
	// 	http.NotFound(w, r)
	// })

	// Wrap the ServeMux with the logging middleware
	// handler := loggingMiddleware(api.mux)

	// Start the server
	// port := 8080
	api.Serve(8080)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
