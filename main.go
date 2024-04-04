package main

import (
	// g"fmt"
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/dgunzy/capstone-webserver/modelApi"
	"github.com/dgunzy/capstone-webserver/routing"
)

func main() {

	port := os.Getenv("PORT")
	router := routing.SetupRoutes()
	fmt.Printf("Server running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

}
