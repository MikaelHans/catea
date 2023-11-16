package login

import (
	"database/sql"

	"github.com/MikaelHans/catea/user-service/api"
	// "github.com/MikaelHans/catea/user-service/pkg/structs"
	"github.com/MikaelHans/catea/user-service/pkg/util"
)

func GetMemberWithLoginInfo(logininfo *api.LoginCredentials) (*sql.Rows, error) {
	db, err := util.GetDBConnection()

	if err != nil {
		return nil, err
	}

	var query string = `SELECT email, password,first_name,last_name,member_since
	FROM catea_member
	WHERE email = $1`

	rows, err := db.Query(query, logininfo.Email)
	db.Close()
	return rows, err
}
