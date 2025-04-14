package main

import (
	"encoding/json"
	"golang_rest_api/internal/config"
	"golang_rest_api/internal/dependencyinjection"
	"golang_rest_api/internal/server"
	"net/http"
	"strconv"
)

func main() {
	// Read config
	config.GetViperConfig()
	// Get dependency injection
	dependencyInjection := dependencyinjection.NewDependencyInjection()

	// Start server
	myServer := server.StartServer()
	myServer.Get("/calc", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		input := query.Get("input")
		intInput, err := strconv.Atoi(input)
		if err != nil {
			message := map[string]string{
				"message": "input is not an integer",
			}
			b1, _ := json.Marshal(message)
			_, _ = w.Write(b1)
		}
		resp := dependencyInjection.RestApi.Calc(intInput)
		_, _ = w.Write(resp)
	})
	server.ListenAndServe(myServer)
}
