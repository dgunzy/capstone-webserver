package routing

import (
	"html/template"
	"net/http"

	"github.com/dgunzy/capstone-webserver/modelApi"
)

// HomeHandler handles requests to the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("routing/templates/index.gohtml"))
	// data := map[string]interface{}{
	// 	"Title": "HTMX Example",
	// }
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SummaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing the form", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("routing/templates/summaryTemplate.gohtml"))

	text := r.FormValue("dynamicTextarea")

	if text == "" {
		return
	}

	// fmt.Println("The input text is: " + text)
	summaryText := modelApi.ModelCaller(text, 1200)
	// fmt.Println("The model produced this summary: " + summaryText)
	if summaryText == "Service Unavailable" {
		if err := tmpl.Execute(w, "Service is down, please wait a few minutes for it to boot up."); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := tmpl.Execute(w, summaryText); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func Texthandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("routing/templates/uploadtext.gohtml"))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Filehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("routing/templates/uploadfile.gohtml"))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing the form", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("routing/templates/summaryTemplate.gohtml"))

	text := r.FormValue("extractedText")

	// fmt.Println("The input text is: " + text)
	summaryText := modelApi.ModelCaller(text, 1200)
	// fmt.Println("The model produced this summary: " + summaryText)
	if summaryText == "Service Unavailable" {
		if err := tmpl.Execute(w, "Service is down, please wait a few minutes for it to boot up."); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if err := tmpl.Execute(w, summaryText); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("routing/templates/about.gohtml"))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
