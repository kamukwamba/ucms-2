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

	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/aboutus", routes.AboutUs)
	http.HandleFunc("/programs", routes.Programs)
	http.HandleFunc("/login", routes.LoginPage)
	http.HandleFunc("/enroll", routes.Enrollment)
	http.HandleFunc("/confirmenrrol", routes.ConfirmEnrollment)
	http.HandleFunc("/adminlogin", routes.AdminLogin)
	http.HandleFunc("/admindashboard", routes.AdminDashboard)
	http.HandleFunc("/programcards", routes.Programcards)

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.ListenAndServe(":3000", nil)

}