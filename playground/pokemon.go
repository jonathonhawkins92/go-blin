package main

import (
	"net/http"

	g "goblin"
	res "goblin/responses"
)

func SomeNewExample() {
	api := GetAPI()
	api.Get(
		g.Config{Path: "/peter"},
		func(request *http.Request, response *res.Response) (*res.Response, error) {
			return response.Accepted().Text("KekW"), nil
		})
}
