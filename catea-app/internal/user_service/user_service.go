package user_service

import (
	"flag"

	user "github.com/MikaelHans/catea/user/pkg/structs"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func Login(email string, pass string){
	var loginInfo user.LoginInfo
	loginInfo.Email = email
	loginInfo.Pass = pass


	
}