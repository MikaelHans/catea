package login

import (
	"fmt"
	"github.com/MikaelHans/catea/user/pkg/structs"
	"github.com/MikaelHans/catea/user/pkg/util"
	// "github.com/MikaelHans/catea/session-management/cmd/session"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



func Login(c *gin.Context) (*structs.LoginResponse){
	var logininfo structs.LoginInfo;
	//CHECK FOR BINDING ERROR///////////////////////////////////////////
	// if err := c.ShouldBindJSON(&logininfo); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	// 	return
	// }
	//QUERY TO DB THEN CHECK FOR QUERY ERROR///////////////////////////////////////////
	rows, err := GetMemberWithLoginInfo(logininfo)
	if err != nil{
		return &structs.LoginResponse{Token: "", Error: err}
	}
	/*
	CHECK IF THE CREDENTIALS(EMAIL, PASS) IS CORRECT BY COUNTING THE ROWS, 0 MEANS INCORRECT SINCE IT MEANS EITHER
	EMAIL OR PASS IS INCORRECT
	*/
	var i int
	var member_data structs.Member;

	for rows.Next(){
		fmt.Print(rows.Scan(
			&member_data.Email,
			&member_data.Pass,
			&member_data.Firstname,
			&member_data.Lastname,
			&member_data.Member_Since))		
		i++
	}
	//VERIFY PASS /////////////////////////////////////////////////////////////////////////////////////////////////
	err = util.DecryptString(member_data.Pass, logininfo.Pass)
	//RETURN 401 CREDENTIALS ARE WRONG ///////////////////////////////////////////////////////////////////////////
	if i < 0 || err != nil{
		var response structs.LoginResponse;
		response.Token = ""
		response.Error = status.Error(codes.Unauthenticated, "Invalid username or password")
		return &structs.LoginResponse{Token: "", Error: status.Error(codes.Unauthenticated, "Invalid username or password")}
	}
	/*WHEN SUCCESS RETURN TOKEN AND MSG:SUCCESS*/
	token, err := util.GenerateJWT(member_data.Email)

	/*GENERATE JWT ERROR HANDLER*/
	if err != nil{
		return &structs.LoginResponse{Token: "", Error: err}
	}

	// responseData := map[string]interface{}{
	// 	"msg": "success",
	// 	"token":token,
	// }
	// c.JSON(http.StatusAccepted, responseData)
	//INITIATE REDIS SESSION////////////////////////////////////////////////////
	// session.StoreSessionData(member_data, token, c)

	return &structs.LoginResponse{Token: token, Error: err}
}