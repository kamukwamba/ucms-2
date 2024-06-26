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

	studentsdata :=
		`create table if not exists studentdata( 
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
													method_of_incounter text);`

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

	fmt.Println("Server running")

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
