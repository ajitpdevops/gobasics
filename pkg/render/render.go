package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Printf("Error parsing template %s", err)
		return
	}
}