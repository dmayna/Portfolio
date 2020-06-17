package main

import (
	"fmt"
	"net/http"
	"uri"
)

//Go application entrypoint
func main() {
	fmt.Println("Listening on port " + uri.PORT)

	router := uri.NewRouter()
	fmt.Println(http.ListenAndServe(uri.PORT, router))
}
