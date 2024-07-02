package routes

import (
	"database/sql/driver"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"ucmps/dbcode"
	"ucmps/encription"
)

type StudentInfo struct {
	UUID                  string
	First_Name            string
	Last_Name             string
	Phone                 string
	Email                 string
	Date_Of_Birth         string
	Gender                string
	Marital_Status        string
	Country               string
	Education_Background  string
	Program               string
	High_School           string
	Grammer_Confirmation  string
	Waiver                string
	Children              string
	School_Attended       string
	Major_In              string
	Degree_Obtained       string
	Current_Occupation    string
	Field_Interested      string
	Prio_Techniques       string
	Previouse_Experience  string
	Purpose_Of_Enrollment string
	Use_Of_Knowledge      string
	Reason_For_Choice     string
	Method_Of_Encounter   string
}

type ACAMS struct {
	UUID                   string
	Student_UUID           string
	First_Name             string
	Last_Name              string
	Email                  string
	Payment_Method         string
	Paid                   string
	Accepted               string
	Communication          string
	Public_Speaking        string
	Intuition              string
	Understanding_Religion string
	Public_Relation        string
	Anger_Management       string
	Connecting_With_Angels string
	Critical_Thinking      string
	Student_Results        string
	Complete               string
}

type StudentCridentials struct {
	UUID        string
	StudentUUID string
	Email       string
	Password    string
}

func SendEMAIL() {
	fmt.Println("New Student Registered")
}

func Validation(email string) bool {

	result := false

	return result
}

func CreateStudentCridentials(studentdate StudentCridentials) bool {

	confirm_creation := true
	dbread := dbcode.SqlRead()
	cridentials, err := dbread.DB.Begin()
	if err != nil {
		log.Fatal()
	}

	stmt, err := cridentials.Prepare("insert into studentcridentials(uuid, student_uuid, email,password) values(?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(studentdate.UUID, studentdate.StudentUUID, studentdate.Email, studentdate.Password)

	if err != nil {
		log.Fatal(err)
		fmt.Println("PART 2: Failed to save to cridentials")
		confirm_creation = false
	}

	err = cridentials.Commit()
	if err != nil {
		log.Fatal(err)
		confirm_creation = false
	}

	return confirm_creation
}

func AddStudentACAMS(data ACAMS) bool {
	dbread := dbcode.SqlRead()
	studentuuid := encription.Generateuudi()
	entry, err := dbread.DB.Begin()
	var result bool = true
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := entry.Prepare("insert into acams(uuid, st_uuid, first_name, last_name, email,payment_type,paid, accepted, communication, public_speaking, intuition, understanding_religion, public_relation,anger_management,connecting_with_angles, critical_thinking,complete) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(studentuuid, data.Student_UUID, data.First_Name, data.Last_Name, data.Email, data.Payment_Method, data.Paid, data.Accepted, data.Communication, data.Public_Speaking, data.Intuition, data.Understanding_Religion, data.Public_Relation, data.Anger_Management, data.Connecting_With_Angels, data.Critical_Thinking, data.Complete)
	if err != nil {
		log.Fatal(err)
		fmt.Println("PART 2: Failed to save acams")
		result = false
	}

	err = entry.Commit()
	if err != nil {
		log.Fatal(err)
		result = false
	}

	return result
}

func CreateStudent(data StudentInfo) bool {
	dbread := dbcode.SqlRead()
	entry, err := dbread.DB.Begin()
	var result bool = true
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := entry.Prepare("insert into studentdata(uuid, first_name, last_name, phone, email, date_of_birth,marital_status,country,eduction_background,program,high_scholl_confirmation,grammer_comprihention,waiver,number_of_children,school_atteneded,major_studied,degree_obtained,current_occupetion,field_interested_in,mps_techqnique_Practiced,previouse_experince,purpose_of_enrollment,use_of_degree,reason_for_choice,method_of_incounter) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.UUID, data.First_Name, data.Last_Name, data.Phone, data.Email, data.Date_Of_Birth, data.Marital_Status, data.Country, data.Education_Background, data.Program, data.High_School, data.Grammer_Confirmation, data.Waiver, data.Children, data.School_Attended, data.Major_In, data.Degree_Obtained, data.Current_Occupation, data.Field_Interested, data.Prio_Techniques, data.Previouse_Experience, data.Purpose_Of_Enrollment, data.Use_Of_Knowledge, data.Reason_For_Choice, data.Method_Of_Encounter)
	if err != nil {
		log.Fatal(err)
		fmt.Println("PART 2: Failed to execute")
		result = false
	}

	err = entry.Commit()
	if err != nil {
		log.Fatal(err)
		result = false
	}

	return result

}

func FindStudent(email string) bool {
	dbread := dbcode.SqlRead()
	var result bool = true

	stmt, err := dbread.DB.Prepare("select first_name from studentdata where email = ?")
	if err != nil {
		log.Fatal(err)
		result = false
	}
	defer stmt.Close()

	var first_name string
	err = stmt.QueryRow(email).Scan(&first_name)

	if err != nil {
		result = false

	}

	return result
}

func Enrollment(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	if r.Method == "POST" {
		fmt.Println("Form obtained")
	}
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	//debug failure to laod templates

	err := tpl.ExecuteTemplate(w, "enrollstudent.html", nil)

	if err != nil {
		log.Fatal(err)
	}
}

type ProgramListName struct {
	Program_Name []string
}

type StringSlice []string

func (stringSlice StringSlice) Value() (driver.Value, error) {
	var quotedStrings []string
	for _, str := range stringSlice {
		quotedStrings = append(quotedStrings, strconv.Quote(str))
	}
	value := fmt.Sprintf("[%s]", strings.Join(quotedStrings, ","))
	return value, nil
}

func AddStudentPrograms(studentuuid, programname string) {
	var programlistname StringSlice
	uuid := encription.Generateuudi()

	programlistname = append(programlistname, programname)

	dbread := dbcode.SqlRead()
	program_name_list, err := dbread.DB.Begin()
	if err != nil {
		log.Fatal()
	}

	stmt, err := program_name_list.Prepare("insert into studentprogramlist(uuid, student_uuid, program_list) values(?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(uuid, studentuuid, programlistname)

	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to create Stduent program list")
	}

	err = program_name_list.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("STUDENT PROGRAM LIST CREATED SUCCESFULLY")

}

func ConfirmEnrollment(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	r.ParseForm()

	var studentsdatain StudentInfo
	uuid := encription.Generateuudi()
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	phone := r.FormValue("phone_number")
	email := r.FormValue("email")
	dateofbirth := r.FormValue("date_of_birth")
	gender := r.FormValue("gender")
	maritalstatus := r.FormValue("marital_status")
	country := r.FormValue("country")
	educationlevel := r.FormValue("education_level")
	program := r.FormValue("program")
	confirmdiplomer := r.FormValue("ucms_diplomer")
	languagecomprihension := r.FormValue("launguage_comprihension")
	waiver := r.FormValue("waiver")
	chidrencount := r.FormValue("children_count")
	collegename := r.FormValue("trade_college_name")
	collegemajor := r.FormValue("college_major")
	collegediplomer := r.FormValue("college_diplomers")
	currentoccupation := r.FormValue("current_occupation")
	fieldofinterest := r.FormValue("field_of_interest")
	priorexperience := r.FormValue("prior_experience")
	priorknowledge := r.FormValue("prior_knowledge")
	purposeofenrolling := r.FormValue("purpose_of_enrolling")
	applicationofknowledge := r.FormValue("application_of_knowledge")
	reasonforchossingucms := r.FormValue("reason_for_chossing_ucms")
	methodofknowledge := r.FormValue("method_of_knowledge")

	studentsdatain = StudentInfo{
		UUID:                  uuid,
		First_Name:            first_name,
		Last_Name:             last_name,
		Phone:                 phone,
		Email:                 email,
		Date_Of_Birth:         dateofbirth,
		Gender:                gender,
		Marital_Status:        maritalstatus,
		Country:               country,
		Education_Background:  educationlevel,
		Program:               program,
		High_School:           confirmdiplomer,
		Grammer_Confirmation:  languagecomprihension,
		Waiver:                waiver,
		Children:              chidrencount,
		School_Attended:       collegename,
		Major_In:              collegemajor,
		Degree_Obtained:       collegediplomer,
		Current_Occupation:    currentoccupation,
		Field_Interested:      fieldofinterest,
		Prio_Techniques:       priorexperience,
		Previouse_Experience:  priorknowledge,
		Purpose_Of_Enrollment: purposeofenrolling,
		Use_Of_Knowledge:      applicationofknowledge,
		Reason_For_Choice:     reasonforchossingucms,
		Method_Of_Encounter:   methodofknowledge,
	}

	chaeck_user := FindStudent(email)

	if chaeck_user {
		fmt.Println("User With eamil already in database")
	} else {
		fmt.Println("Create New User")
		result := CreateStudent(studentsdatain)
		if result {
			stuuuid := encription.Generateuudi()
			paid := "complete"
			payment_method := "lamp"
			accepted := "false"
			communication := "incomplete"
			public_speaking := "incomplete"
			intuition := "incomplete"
			understanding_religion := "incomplete"
			public_relation := "incomplete"
			anger_management := "incomplete"
			connecting_with_angles := "incomplete"
			critical_thinking := "incomplete"
			complete := "incomplete"

			addstudentacams := ACAMS{
				UUID:                   stuuuid,
				Student_UUID:           uuid,
				First_Name:             first_name,
				Last_Name:              last_name,
				Email:                  email,
				Payment_Method:         payment_method,
				Paid:                   paid,
				Accepted:               accepted,
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
			addedtoacams := AddStudentACAMS(addstudentacams)

			if addedtoacams {
				SendEMAIL()
				cridentuuid := encription.Generateuudi()
				studentcridentials := StudentCridentials{
					UUID:        cridentuuid,
					StudentUUID: uuid,
					Email:       email,
					Password:    email,
				}
				CreateStudentCridentials(studentcridentials)
				AddStudentPrograms(uuid, "ACAMS")
			} else {
				fmt.Println("Problem with adding student to acams")
			}
		} else {
			fmt.Println("FAILED TO CREATE NEW USER")
		}

		err := tpl.ExecuteTemplate(w, "confirmenroll.html", nil)

		if err != nil {
			log.Fatal(err)
		}
	}

}
