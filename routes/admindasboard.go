package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ucmps/dbcode"
)

type AdminLogData struct {
	Email    string
	Password string
}
type AdminInfo struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func AdminAuth(data AdminLogData, dataList []dbcode.AdminInfo) (bool, AdminInfo) {

	var result bool
	var admin_data AdminInfo
	for _, admin_info := range dataList {
		id := admin_info.ID
		name := admin_info.Name
		email := admin_info.Email
		password := admin_info.Password

		if data.Password == password && data.Email == email {
			fmt.Println("A match was found")
			admin_data = AdminInfo{
				ID:       id,
				Name:     name,
				Email:    email,
				Password: password,
			}
			result = true
		}
	}
	return result, admin_data
}

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	if r.Method == "POST" {
		r.ParseForm()

		adminList := dbcode.AdminGet()

		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		authget := AdminLogData{
			Email:    email,
			Password: password,
		}

		check, admin_dataout := AdminAuth(authget, adminList)

		if check {
			fmt.Println("redirecting")
			err := tpl.ExecuteTemplate(w, "admindasboard.html", admin_dataout)

			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Check email or user name")
		}

	} else {
		fmt.Fprint(w, "Permision is required for this")
	}
	// err := tpl.ExecuteTemplate(w, "admindasboard.html", nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

}
