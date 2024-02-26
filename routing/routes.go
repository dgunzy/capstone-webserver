package routing

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("routing/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/summarize", SummaryHandler)
	mux.HandleFunc("/uploadtext", Texthandler)
	mux.HandleFunc("/uploadfile", Filehandler)
	return mux
}
