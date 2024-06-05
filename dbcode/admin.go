package dbcode

import (
	"database/sql"
	"fmt"
	"log"
)

type AdminInfo struct {
	ID       string
	Name     string
	Email    string
	Password string
}

var inforOutLsit []AdminInfo

func AdminGet() []AdminInfo {

	dbread := SqlRead()
	var infor_out AdminInfo
	rows, err := dbread.DB.Query("select uuid, admin_name, admin_email, admin_password from admin")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var admin_name string
		var admin_email string
		var admin_password sql.NullString
		err = rows.Scan(&id, &admin_name, &admin_email, &admin_password)

		infor_out = AdminInfo{
			ID:       id,
			Name:     admin_name,
			Email:    admin_email,
			Password: admin_password.String,
		}
		inforOutLsit = append(inforOutLsit, infor_out)
		if err != nil {
			log.Fatal(err)
		}

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return inforOutLsit
}

func CreateAdmin(info AdminInfo) {
	dbread := SqlRead()
	tx, err := dbread.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into admin(uuid, admin_name, admin_email, admin_password) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
