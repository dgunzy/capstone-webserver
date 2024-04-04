package main

import (
	// g"fmt"
	"fmt"
	"log"
	"net/http"

	// "github.com/dgunzy/capstone-webserver/modelApi"
	"github.com/dgunzy/capstone-webserver/routing"
)

func main() {

	router := routing.SetupRoutes()
	fmt.Println("Server running on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
