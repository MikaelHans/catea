package signup

import (
	"database/sql"
	"github.com/MikaelHans/catea/login-signup/pkg/util"
	"github.com/MikaelHans/catea/login-signup/pkg/structs"
)

func InsertIntoMember(newMember structs.Member) (*sql.Rows, error) {
	db, err := util.GetDBConnection()

	if err != nil {
		return nil, err
	}

	var query string = `SELECT insert_member($1,$2,$3,$4) as result`

	rows, err := db.Query(query, newMember.Email, newMember.Pass, newMember.Firstname, newMember.Lastname)
	db.Close()
	return rows, err
}


