package util

import (
	"database/sql"
	"log"
)

func GetMemberWithEmail(email string)(*sql.Rows, error){
	db, err := GetDBConnection()
	
	if (err != nil){
		log.Fatal(err)
	}

	var query string = `
	SELECT *
	FROM catea_member
	WHERE email = '$1'`

	rows, err := db.Query(query, email)
	return rows, err
}