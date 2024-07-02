package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ucmps/dbcode"
)

type ACMS struct {
	UUID                           string
	Student_UUID                   string
	First_Name                     string
	Last_Name                      string
	Email                          string
	Payment_Method                 string
	Paid                           string
	Accepted                       string
	Student_Results                string
	Complete                       string
	Mindfulness                    string
	Dreams_and_Dreaming            string
	Energy_of_Money                string
	Crystals_and_Gemstones         string
	Forgiveness                    string
	Cleansing_and_Fasting          string
	Astrology                      string
	African_Culture_and_Traditions string
	Transforming_personalities     string
}

type ADMS struct {
	UUID                                    string
	Student_UUID                            string
	First_Name                              string
	Last_Name                               string
	Email                                   string
	Payment_Method                          string
	Paid                                    string
	Accepted                                string
	Student_Results                         string
	Complete                                string
	Creative_Writing                        string
	Understanding_Miracles                  string
	Channeling_skills                       string
	Enneagram                               string
	Mythology_on_Gods_and_Goddess           string
	Herbs                                   string
	Meditation_skills                       string
	Mantras_and_Mudras                      string
	Divinations                             string
	Archetypes                              string
	Basics_in_Research                      string
	Understanding_Propaganda                string
	Great_Spiritual_Teachers                string
	Reprogramming                           string
	Shamanism                               string
	Mystery_Schools_in_the_world            string
	Law_and_Ethics_in_Metaphysical_Sciences string
	Non_Violet_Communication                string
}

type ABDMS struct {
	UUID                             string
	Student_UUID                     string
	First_Name                       string
	Last_Name                        string
	Email                            string
	Payment_Method                   string
	Paid                             string
	Accepted                         string
	Student_Results                  string
	Complete                         string
	Cause_and_Core_Issues_in_Beliefs string
	Emotional_Well_Being             string
	The_Art_of_Breathing             string
	Spiritual_symbols_and_colours    string
	Psychic_Skills                   string
	Shadow_Work                      string
	The_Craft                        string
	Hypnosis_and_Beyond              string
	Mysterious_experiences           string
	Manifestation_skills             string
	Unlocking_Creativity             string
	Transpersonal_counselling        string
	African_Healing_Arts             string
	Ceremonies_of_the_World          string
	Mother_Earth                     string
	The_Art_of_Placement             string
	Chakras_and_Auras                string
	Transforming_personalities       string
	Mayan_Calendar                   string
	Polarity_Therapy                 string
	Introduction_To_Meditation       string
	Health_and_Nutrition             string
	Setting_up_a_business            string
}

type StudentCourse struct {
	ACAMSCourse ACAMS
	ACMSCourse  ACMS
	ADMSCourse  ADMS
	ABDMSCourse ABDMS
}

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

func ConfirmStudentLogin(w http.ResponseWriter, r *http.Request) {

	var students_data_acams ACAMS

	tpl = template.Must(template.ParseGlob("templates/*.html"))

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
			studentprogramlist := GetStudentPrograms(studentuuid)

			fmt.Println("Programs Student Has Been Accepted For: ", studentprogramlist)
		}

	}

	students_data := StudentCourse{
		ACAMSCourse: students_data_acams,
	}
	fmt.Println(students_data)

	tpl.ExecuteTemplate(w, "studentportal.html", students_data)
}
