package main

import (
	"fmt"
	"net/http"
	"Portfolio/src/uri"
	"os"
)

//Go application entrypoint
func main() {
	port :=  ":" + os.Getenv("PORT")
	fmt.Println("Listening on port " + port)

	router := uri.NewRouter()
	fmt.Println(http.ListenAndServe(port, router))
}
