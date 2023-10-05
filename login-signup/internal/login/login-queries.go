package login
import (
	"database/sql"
	"github.com/MikaelHans/catea/login-signup/pkg/util"
	"github.com/MikaelHans/catea/login-signup/pkg/structs"
)
func GetMemberWithLoginInfo(logininfo structs.LoginInfo) (*sql.Rows, error) {
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