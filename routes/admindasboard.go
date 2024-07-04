package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
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

type DashData struct {
	ACAMSTotal int
	AdminInfo
}

type StudentProgramList struct {
	Program_Name string
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
		acamscount := ACAMSCount()
		adminList := dbcode.AdminGet()

		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		authget := AdminLogData{
			Email:    email,
			Password: password,
		}

		check, admin_dataout := AdminAuth(authget, adminList)

		toshow := DashData{
			ACAMSTotal: acamscount,
			AdminInfo:  admin_dataout,
		}

		if check {
			fmt.Println("redirecting")
			err := tpl.ExecuteTemplate(w, "admindasboard.html", toshow)

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

func StudentACAMSData(student_uuid string) {
	dbread := dbcode.SqlRead()

	stmt, err := dbread.DB.Prepare("select program_list from studentprogramlist where student_uuid = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var program_list string

	err = stmt.QueryRow(student_uuid).Scan(&program_list)

	if err != nil {
		fmt.Println("FAILED TO GET STUDENT PROGRAM LIST")
	}

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
			Student_UUID:   student_uuid,
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

func ACAMSCount() int {
	dbread := dbcode.SqlRead()
	var counter int

	rows, err := dbread.DB.Query("select * from  acams")
	if err != nil {
		fmt.Println("Failed to get acams student data")
	}
	defer rows.Close()

	for rows.Next() {
		counter += 1

	}

	return counter
}

func GetStudentPrograms(student_uuid string) []string {
	dbread := dbcode.SqlRead()

	stmt, err := dbread.DB.Prepare("select program_list from studentprogramlist where student_uuid = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var program_list string

	err = stmt.QueryRow(student_uuid).Scan(&program_list)

	trimedlist := strings.Trim(program_list, "[]\"")
	listout := strings.Split(trimedlist, ",")

	if err != nil {
		fmt.Println("FAILED TO GET STUDENT PROGRAM LIST")
	}

	return listout
}

func GetStudentAllDetails(uuid string) StudentInfo {
	dbread := dbcode.SqlRead()

	stmt, err := dbread.DB.Prepare("select  uuid, first_name, last_name, phone, email, date_of_birth,marital_status,country,eduction_background,program,high_scholl_confirmation,grammer_comprihention,waiver,number_of_children,school_atteneded,major_studied,degree_obtained,current_occupetion,field_interested_in,mps_techqnique_Practiced,previouse_experince,purpose_of_enrollment,use_of_degree,reason_for_choice,method_of_incounter from studentdata where uuid = ?")

	if err != nil {
		log.Fatal(err)

	}
	defer stmt.Close()

	var data StudentInfo

	err = stmt.QueryRow(uuid).Scan(&data.UUID, &data.First_Name, &data.Last_Name, &data.Phone, &data.Email, &data.Date_Of_Birth, &data.Marital_Status, &data.Country, &data.Education_Background, &data.Program, &data.High_School, &data.Grammer_Confirmation, &data.Waiver, &data.Children, &data.School_Attended, &data.Major_In, &data.Degree_Obtained, &data.Current_Occupation, &data.Field_Interested, &data.Prio_Techniques, &data.Previouse_Experience, &data.Purpose_Of_Enrollment, &data.Use_Of_Knowledge, &data.Reason_For_Choice, &data.Method_Of_Encounter)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("student info: ", data)

	return data
}

//ROUTER CODE

func StudentProfileData(w http.ResponseWriter, r *http.Request) {

	studentuuid := r.PathValue("id")
	studentdataout := GetStudentAllDetails(studentuuid)
	listout := GetStudentPrograms(studentuuid)

	fmt.Println("Student UUID: ", studentuuid)
	fmt.Println("Student Data: ", studentdataout)
	fmt.Println("Student Program List: ", listout)

	for _, program := range listout {
		fmt.Println("Program Name: ", program)
		switch program {
		case "ACAMS":
			StudentACAMSData(studentuuid)
			fmt.Println("Certificate Program")
		}

	}

	tpl = template.Must(template.ParseGlob("templates/*.html"))

	err := tpl.ExecuteTemplate(w, "studentdetailstemplate.html", studentdataout)

	if err != nil {
		log.Fatal(err)
	}

}

func ACAMSStudentData(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	acamsstudents := GetACAMSStudents()

	err := tpl.ExecuteTemplate(w, "studentdataadmin.html", acamsstudents)

	if err != nil {
		log.Fatal(err)
	}
}
