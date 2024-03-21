package routing

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"github.com/dgunzy/capstone-webserver/modelApi"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("routing/templates/index.gohtml"))

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
	reg, err := regexp.Compile(`["']|\s{2,}`)
	if err != nil {
		fmt.Println("Error compiling regex ", err)
		return
	}

	newText := reg.ReplaceAllString(text, "")

	// fmt.Println("The input text is: " + text)
	summaryText := modelApi.ModelCaller(newText, 5000)
	// fmt.Println("The model produced this summary: " + summaryText)
	if summaryText == "Service Unavailable" {
		if err := tmpl.Execute(w, "AI Endpoint is down, please wait a 2 - 5 minutes for it to boot up for your session."); err != nil {
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

	reg, err := regexp.Compile(`["']|\s{2,}`)
	if err != nil {
		fmt.Println("Error compiling regex ", err)
		return
	}

	text = reg.ReplaceAllString(text, "")

	// fmt.Println("The input text is: " + text)
	summaryText := modelApi.ModelCaller(text, 5000)
	// fmt.Println("The model produced this summary: " + summaryText)
	if summaryText == "Service Unavailable" {
		if err := tmpl.Execute(w, "AI Endpoint is down, please wait a 2 - 5 minutes for it to boot up for your session."); err != nil {
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

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("routing/templates/contact.gohtml"))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
