package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func ValidateSudent() bool {
	isstudent := true

	return isstudent

}

func ConfirmStudentLogin(w http.ResponseWriter, r *http.Request) {

	tpl = template.Must(template.ParseGlob("templates/*.html"))

	r.ParseForm()
	if r.Method == "POST" {
		fmt.Println("form is obtained")
		studentemail := r.FormValue("studentemail")
		studentpassword := r.FormValue("studentpassword")

		fmt.Println("Student Name: ", studentemail, "Student Password: ", studentpassword)

	}

	tpl.ExecuteTemplate(w, "studentportal.html", nil)
}
