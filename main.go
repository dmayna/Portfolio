package main

import (
	"fmt"
	"net/http"
	"Portfolio/src/uri"
	
)

//Go application entrypoint
func main() {
	port :=  ":8080"// + os.Getenv("PORT")
	fmt.Println("Listening on port " + port)

	router := uri.NewRouter()
	fmt.Println(http.ListenAndServe(port, router))
}
