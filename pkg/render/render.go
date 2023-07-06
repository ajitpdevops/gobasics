package render

import (
	"html/template"
	"log"
	"net/http"
)

// TODO:
// 1. Do not load template to disk every time
// 2. Create a template cache
// 3. Check if the data already exists in the templateCache
// 4. Reload only if it does not exist and save the data in the templateCache

func RenderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Printf("Error parsing template %s", err)
		return
	}
}