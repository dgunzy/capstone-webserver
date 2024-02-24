package routing

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/summarize", SummaryHandler)
	return mux
}
