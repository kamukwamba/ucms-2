package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func AboutUs(w http.ResponseWriter, r *http.Request) {

	tpl = template.Must(template.ParseGlob("templates/*.html"))

	//debug failure to laod templates
	fmt.Println("Working")
	err := tpl.ExecuteTemplate(w, "aboutus.html", nil)
	fmt.Println("Working")

	if err != nil {
		log.Fatal(err)
	}
}
