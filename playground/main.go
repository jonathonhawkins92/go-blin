package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"fire"
	g "goblin"
	"goblin/responses"
	resesesponse "goblin/responses"

	"github.com/getkin/kin-openapi/openapi3"
)

//go:generate go run goblin/analyze

var API *g.Client

type ImportedExample struct {
	fire.APIWrapper
}

type APIWrapperPreChildChild struct {
	api APIWrapperPreChild
}

type APIWrapperPreChild struct {
	api APIWrapperPrePre
}

type APIWrapperPrePre struct {
	APIWrapperPre
}

type APIWrapperPre struct {
	APIWrapper
}
type APIWrapper struct {
	g.Client
}

type APIWrapperPost struct {
	APIWrapper
}

type APIWrapperPostPost struct {
	APIWrapperPost
}

type APIWrapperPostChild struct {
	api APIWrapperPostPost
}

type APIWrapperPostChildChild struct {
	api APIWrapperPostChild
}

// type SomeResponseHelper struct {
// 	resesesponse.Response
// }

func (a *APIWrapper) Thing() {
	a.Get(g.Config{
		Path: "/example",
	}, Endpointerino)
}

func GetAPI() *g.Client {
	if API != nil {
		return API
	}

	return g.API(openapi3.Info{
		Title:       "My example API",
		Version:     "0.0.1",
		Description: "My cool app",
	})
}

func Endpointerino(request *http.Request, reeeeesponse *resesesponse.Response) (*resesesponse.Response, error) {
	pokemon, err := fetchPokemonData("pikachu")
	if err != nil {
		return reeeeesponse.Http500().Text("bang").BadGateway().Accepted().AlreadyReported().Created().Http102().Http204NoContent(), nil
	}
	return reeeeesponse.Http200().JSON(pokemon), nil
}

type Other struct {
	Name string `json:"name"`
}

type PokemonResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Height int   `json:"height"`
	Weight int   `json:"weight"`
	Other  Other `json:"other"`
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
	api := GetAPI()

	api.Get(g.Config{
		Path:        "/",
		Summary:     "Something nice",
		Description: "Something super dooper nice",
	}, func(request *http.Request, reeeeesponse *responses.Response) (*resesesponse.Response, error) {
		return reeeeesponse.Http200().Text("hi"), nil
	})

	api.Get(g.Config{
		Path: "/example",
	}, func(request *http.Request, reeeeesponse *resesesponse.Response) (*resesesponse.Response, error) {
		pokemon, err := fetchPokemonData("pikachu")
		if err != nil {
			return reeeeesponse.Http500().Text("bang").BadGateway().Accepted().AlreadyReported().Created().Http102().Http204NoContent(), nil
		}
		return reeeeesponse.Http200().JSON(pokemon), nil
	})

	// api.Put(goblin.Config{
	// 	Path: "/my-orbies",
	// }, func(request *http.Request, response *responses.Response) (*responses.Response, error) {
	// 	return response.Continue().Text("hi"), nil
	// })

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
