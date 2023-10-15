package login

import (
	"database/sql"

	"github.com/MikaelHans/catea/user/api"
	// "github.com/MikaelHans/catea/user/pkg/structs"
	"github.com/MikaelHans/catea/user/pkg/util"
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