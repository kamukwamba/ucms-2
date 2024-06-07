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

type ReturnACAMS struct {
	Counter                int
	UUID                   string
	Student_UUID           string
	First_Name             string
	Last_Name              string
	Email                  string
	Program                string
	Accepted               string
	Paid                   string
	Payment_Method         string
	Communication          string
	Public_Speaking        string
	Intuition              string
	Understanding_Religion string
	Public_Relation        string
	Anger_Management       string
	Connecting_With_Angels string
	Critical_Thinking      string
	Completed              string
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

func GetACAMSStudents() []ReturnACAMS {
	dbread := dbcode.SqlRead()
	var conuter int
	var datalist []ReturnACAMS
	rows, err := dbread.DB.Query("select * from  acams")
	if err != nil {
		fmt.Println("Failed to get acams student data")
	}
	defer rows.Close()

	for rows.Next() {
		conuter += 1
		var uuid string
		var student_uuid string
		var first_name string
		var last_name string
		var email string
		var accepted string
		var paid string
		var payment_method string
		var communication string
		var public_speaking string
		var intuition string
		var understanding_religion string
		var public_relation string
		var anger_management string
		var connecting_with_angles string
		var critical_thinking string
		var complete string

		err := rows.Scan(&uuid, &student_uuid, &first_name, &last_name, &email, &accepted, &payment_method, &paid, &communication, &public_speaking, &intuition, &understanding_religion, &public_relation, &anger_management, &connecting_with_angles, &critical_thinking, &complete)

		if err != nil {
			fmt.Println("Check the scan for student data")
			log.Fatal(err)
		}

		dataout := ReturnACAMS{
			Counter:        conuter,
			UUID:           uuid,
			First_Name:     first_name,
			Last_Name:      last_name,
			Email:          email,
			Accepted:       accepted,
			Program:        "ACAMS",
			Payment_Method: payment_method,
			Paid:           paid,
			Completed:      complete,
		}

		datalist = append(datalist, dataout)

	}

	return datalist

}

func ACAMSStudentData(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	acamsstudents := GetACAMSStudents()

	err := tpl.ExecuteTemplate(w, "studentdataadmin.html", acamsstudents)

	if err != nil {
		log.Fatal(err)
	}
}
