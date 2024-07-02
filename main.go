package main

import (
	"fmt"
	"log"
	"net/http"
	"ucmps/dbcode"
	"ucmps/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))

	dbread := dbcode.SqlRead()

	defer dbread.DB.Close()

	studentprogramlist :=
		`create table if not exists studentprogramlist(
		uuid blob not null,
		student_uuid blob,
		program_list blob
	)`
	studentcridentials :=
		`create table if not exists studentcridentials(
		uuid blob not null,
		student_uuid blob,
		email text,
		password text
	)`

	sqlacms := `
		create table if not exists acms(
			uuid blob not null,
			student_uuid not null,
			first_name text,
			last_name text,
			email text,
			payment_type text,
			paid text,
			accepted text,
			Mindfulness text,
			Dreams_and_Dreaming text,
			Energy_of_Money text,
			Crystals_and_Gemstones text,
			Forgiveness text,
			Cleansing_and_Fasting text,
			Astrology text,
			African_Culture_and_Traditions text,
			Transforming_personalities text,
			complete text);
	`

	sqlabdms := `create table if not exists abdms(
		uuid_blob_not_null,
		Causes_and_Core_Issues_in_Beliefs text,
		Emotional_Well_Being text,
		The_Art_of_Breathing text,
		Spiritual_symbols_and_colours text,
		Psychic_Skills text,
		Shadow_Work text,
		The_Craft text,
		Hypnosis_and_Beyond text,
		Mysterious_experiences text,
		Manifestation_skills text,
		Unlocking_Creativity text,
		Transpersonal_counselling text,
		African_Healing_Arts text,
		Ceremonies_of_the_World text,
		Mother_Earth text,
		The_Art_of_Placement text,
		Chakras_and_Auras text,
		Transforming_personalities text,
		Mayan_Calendar text,
		Polarity_Therapy text,
		Introduction_To_Meditation text,
		Health_and_Nutrition text,
		Setting_up_a_business text
		);`

	sqladms := `
		create table if not exists adms(
			uuid blob not null,
			student_uuid,
			Creative_Writing text,
			Understanding_Miracles text,
			Channeling_skills text,
			Enneagram text,
			Mythology_on_Gods_and_Goddess text,
			Herbs text,
			Meditation_skills text,
			Mantras_and_Mudras text,
			Divinations text,
			Archetypes text,
			Basics_in_Research text,
			Understanding_Propaganda text,
			Great_Spiritual_Teachers text,
			Reprogramming text,
			Shamanism text,
			Mystery_Schools_in_the_world text,
			Law_and_Ethics_in_Metaphysical_Sciences text,
			Non_Violet_Communication text,
			complete text

		);`

	sqlacams := ` 
		create table if not exists acams(
			uuid blod not null,
			st_uuid blob not null,
			first_name text,
			last_name text,
			email text,
			payment_type text,
			paid text,
			accepted text,
			communication text,
			public_speaking text,
			intuition text,
			understanding_religion text,
			public_relation text,
			anger_management text,
			connecting_with_angles text,
			critical_thinking text,
			complete text);
		`
	sqlStmt := `
		create table if not exists admin(
		uuid blob not null,
		admin_name text, 
		admin_email text, 
		admin_password text);
	`

	studentsdata := `create table if not exists studentdata( 
												uuid blob not null, 
												first_name text, 
												last_name text,
												phone text,
												email text, 
												date_of_birth text, 
												gender text,
												marital_status text, 
												country text, 
												eduction_background text, 
												program text, 
												high_scholl_confirmation text,
												grammer_comprihention text, 
												waiver text, 
												number_of_children text,
												school_atteneded text, 
												major_studied text, 
												degree_obtained text, 
												current_occupetion text,
												field_interested_in text, 
												mps_techqnique_Practiced text, 
												previouse_experince text, 
												purpose_of_enrollment text, 
												use_of_degree text, 
												reason_for_choice text, 
												method_of_incounter text);
		`

	_, errsqladms := dbread.DB.Exec(sqladms)
	if errsqladms != nil {
		log.Printf("%q: %s\n", errsqladms, sqladms)
		return
	}

	_, errsqlabdms := dbread.DB.Exec(sqlabdms)
	if errsqlabdms != nil {
		log.Printf("%q: %s\n", errsqlabdms, sqlabdms)
		return
	}
	_, erracms := dbread.DB.Exec(sqlacms)

	if erracms != nil {
		log.Printf("%q: %s\n", erracms, sqlacms)
		return
	}
	_, errstp := dbread.DB.Exec(studentprogramlist)
	if errstp != nil {
		log.Printf("%q: %s\n", errstp, studentprogramlist)
		return
	}

	_, errstc := dbread.DB.Exec(studentcridentials)
	if errstc != nil {
		log.Printf("%q: %s\n", errstc, studentcridentials)
		return
	}

	_, erracams := dbread.DB.Exec(sqlacams)
	if erracams != nil {
		log.Printf("%q: %s\n", erracams, sqlStmt)
		return
	}

	_, errstd := dbread.DB.Exec(studentsdata)

	if errstd != nil {
		log.Printf("%q: %s\n", errstd, sqlStmt)
		return
	}
	_, err := dbread.DB.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	defer dbread.DB.Close()

	fmt.Println("::SERVER STARTED::")

	router := http.NewServeMux()
	router.HandleFunc("/", routes.HomePage)
	router.HandleFunc("/aboutus/{id}", routes.AboutUs)
	router.HandleFunc("/programs", routes.Programs)
	router.HandleFunc("/login", routes.LoginPage)
	router.HandleFunc("/enroll", routes.Enrollment)
	router.HandleFunc("/confirmenrrol", routes.ConfirmEnrollment)
	router.HandleFunc("/adminlogin", routes.AdminLogin)
	router.HandleFunc("/admindashboard", routes.AdminDashboard)
	router.HandleFunc("/programcards", routes.Programcards)
	router.HandleFunc("/acamsstudentdata", routes.ACAMSStudentData)
	router.HandleFunc("/confirmlogin", routes.ConfirmStudentLogin)
	router.HandleFunc("/studentprofiledata/{id}", routes.StudentProfileData)

	router.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.ListenAndServe(":3000", router)

}
