package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ucmps/dbcode"
)

func ValidateSudent(emailin string) (bool, string) {
	isstudent := true
	dbread := dbcode.SqlRead()
	stmt, err := dbread.DB.Prepare("select uuid, student_uuid, email, password from studentcridentials where email = ?")

	if err != nil {
		isstudent = false
		fmt.Println("First err")
		log.Fatal(err)
	}

	defer stmt.Close()

	var uuid string
	var student_uuid string
	var email string
	var password string

	err = stmt.QueryRow(emailin).Scan(&uuid, &student_uuid, &email, &password)

	if err != nil {
		fmt.Println("Second err")
		log.Fatal(err)
		isstudent = false
	}

	return isstudent, student_uuid

}

func GetFromACAMS(uuidin string) ACAMS {
	fmt.Println(uuidin)

	confirmacams := true
	dbread := dbcode.SqlRead()
	stmt, err := dbread.DB.Prepare("select uuid, st_uuid,first_name, last_name, email, payment_type, paid, accepted, communication, public_speaking, intuition, understanding_religion, public_relation,anger_management,connecting_with_angles,critical_thinking,complete from acams where st_uuid = ?")

	if err != nil {
		fmt.Println("The one not working")
		log.Fatal(err)
	}

	defer stmt.Close()

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

	err = stmt.QueryRow(uuidin).Scan(&uuid, &student_uuid, &first_name, &last_name, &email, &accepted, &payment_method, &paid, &communication, &public_speaking, &intuition, &understanding_religion, &public_relation, &anger_management, &connecting_with_angles, &critical_thinking, &complete)

	if err != nil {
		fmt.Println("Second err")
		log.Fatal(err)
		confirmacams = false
	}

	studentacamsdata := ACAMS{
		UUID:                   uuid,
		Student_UUID:           student_uuid,
		First_Name:             first_name,
		Last_Name:              last_name,
		Email:                  email,
		Accepted:               accepted,
		Payment_Method:         payment_method,
		Paid:                   paid,
		Communication:          communication,
		Public_Speaking:        public_speaking,
		Intuition:              intuition,
		Understanding_Religion: understanding_religion,
		Public_Relation:        public_relation,
		Anger_Management:       anger_management,
		Connecting_With_Angels: connecting_with_angles,
		Critical_Thinking:      critical_thinking,
		Complete:               complete,
	}

	fmt.Println("ACAMS Out: ", studentacamsdata, confirmacams)

	return studentacamsdata
}

type StudentCourse struct {
	ACAMSCourse ACAMS
}

func ConfirmStudentLogin(w http.ResponseWriter, r *http.Request) {

	var students_data_acams ACAMS

	tpl = template.Must(template.ParseGlob("templates/*/*.html"))

	r.ParseForm()
	if r.Method == "POST" {
		fmt.Println("form is obtained")
		studentemail := r.FormValue("studentemail")
		studentpassword := r.FormValue("studentpassword")

		fmt.Println("Student Name: ", studentemail, "Student Password: ", studentpassword)
		idvalue := r.PathValue("id")
		fmt.Println("ID Value", idvalue)
		confirm, studentuuid := ValidateSudent(studentemail)
		if confirm {
			students_data_acams = GetFromACAMS(studentuuid)
			fmt.Println("From ACAMS", students_data_acams)
		}

	}

	students_data := StudentCourse{
		ACAMSCourse: students_data_acams,
	}
	fmt.Println(students_data)

	tpl.ExecuteTemplate(w, "studentportal.html", students_data)
}
