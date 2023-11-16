package signup

import (
	"database/sql"

	"github.com/MikaelHans/catea/user-service/api"
	// "github.com/MikaelHans/catea/user-service/pkg/structs"
	"github.com/MikaelHans/catea/user-service/pkg/util"
)

func InsertIntoMember(newMember *api.SignupCredentials) (*sql.Rows, error) {
	db, err := util.GetDBConnection()

	if err != nil {
		return nil, err
	}

	var query string = `SELECT insert_member($1,$2,$3,$4) as result`

	// var query string = `INSERT INTO catea_member (email, password, first_name, last_name, member_since)
	// VALUES ($1, $2, $3, $4, now())
	// ON CONFLICT (email) DO nothing
	// RETURNING false as result;`

	rows, err := db.Query(query, newMember.Email, newMember.Pass, newMember.Firstname, newMember.Lastname)
	db.Close()
	return rows, err
}
