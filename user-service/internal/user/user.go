package user

import (
	"context"
	"fmt"

	pb "github.com/MikaelHans/catea/user-service/api"
	"github.com/MikaelHans/catea/user-service/internal/user/repository/login"
	"github.com/MikaelHans/catea/user-service/internal/user/repository/signup"
	"github.com/MikaelHans/catea/user-service/internal/user/rpcs/session"
	"github.com/MikaelHans/catea/user-service/pkg/structs"
	"github.com/MikaelHans/catea/user-service/pkg/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) Login(ctx context.Context, logincredentials *pb.LoginCredentials) (*pb.LoginResponse, error) {
	rows, err := login.GetMemberWithLoginInfo(logincredentials)
	if err != nil {
		return &pb.LoginResponse{Token: ""}, status.Error(codes.Internal, err.Error())
	}
	/*
		CHECK IF THE CREDENTIALS(EMAIL, PASS) IS CORRECT BY COUNTING THE ROWS, 0 MEANS INCORRECT SINCE IT MEANS EITHER
		EMAIL OR PASS IS INCORRECT
	*/
	var i int
	var member_data structs.Member

	for rows.Next() {
		fmt.Print(rows.Scan(
			&member_data.Email,
			&member_data.Pass,
			&member_data.Firstname,
			&member_data.Lastname,
			&member_data.Member_Since))
		i++
	}
	//VERIFY PASS /////////////////////////////////////////////////////////////////////////////////////////////////
	err = util.DecryptString(member_data.Pass, logincredentials.Pass)
	//RETURN 401 CREDENTIALS ARE WRONG ///////////////////////////////////////////////////////////////////////////
	if i < 0 || err != nil {
		var response structs.LoginResponse
		response.Token = ""
		response.Error = status.Error(codes.Unauthenticated, "Invalid username or password")
		return &pb.LoginResponse{Token: ""}, status.Error(codes.InvalidArgument, "Invalid username or password")
	}
	/*WHEN SUCCESS RETURN TOKEN AND MSG:SUCCESS*/
	token, err := util.GenerateJWT(member_data.Email)

	/*GENERATE JWT ERROR HANDLER*/
	if err != nil {
		return &pb.LoginResponse{Token: ""}, status.Error(codes.Unauthenticated, "Invalid username or password")
	}

	// responseData := map[string]interface{}{
	// 	"msg": "success",
	// 	"token":token,
	// }
	// c.JSON(http.StatusAccepted, responseData)
	//INITIATE REDIS SESSION////////////////////////////////////////////////////
	// session.StoreSessionData(member_data, token, c)
	data, err := session.SetSession(token, structs.Member{
		Email:        member_data.Email,
		Firstname:    member_data.Firstname,
		Lastname:     member_data.Lastname,
		Member_Since: member_data.Member_Since,
	})

	//variable dump, currently i dont know the best way to handle empty responses from gRPC
	if err != nil {
		return &pb.LoginResponse{Token: data}, err
	}

	return &pb.LoginResponse{Token: token}, nil
}

func (s *Server) SignUp(ctx context.Context, signupcredentials *pb.SignupCredentials) (*pb.SignUpResponse, error) {
	// var data structs.Member
	// if err := c.ShouldBindJSON(&data); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	//ENCRYPT MEMBER PASSWORD//////////////////////////////////////////////////
	encrypted_pass, err := util.EncryptString(signupcredentials.Pass)
	if err != nil {
		return &pb.SignUpResponse{}, status.Error(codes.Internal, err.Error())
	}
	signupcredentials.Pass = encrypted_pass
	//INSERT MEMBER TO DATABASE//////////////////////////////////////////////////
	rows, err := signup.InsertIntoMember(signupcredentials)
	if err != nil {
		return &pb.SignUpResponse{}, status.Error(codes.Internal, err.Error())
	}
	var result bool
	rows.Next()
	err = rows.Scan(&result)
	if err != nil {
		return &pb.SignUpResponse{}, status.Error(codes.Internal, err.Error())
	}
	if result == false {
		return &pb.SignUpResponse{}, status.Error(codes.AlreadyExists, "email already used by another account")
	}
	return &pb.SignUpResponse{Message: "OK"}, nil
}
