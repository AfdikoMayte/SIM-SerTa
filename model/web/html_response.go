package web

import (
	"html/template"
	"net/http"
)

type DataToPage map[string]interface{}

func HtmlResponse(w http.ResponseWriter, path string, namePage string, dataToPage DataToPage) {

	var tmpl = template.Must(template.ParseFiles(
		"views/"+path,
		"views/partial/_header.html",
		"views/partial/_footer.html",
		"views/partial/_navbar.html",
	))

	err := tmpl.ExecuteTemplate(w, namePage, dataToPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
