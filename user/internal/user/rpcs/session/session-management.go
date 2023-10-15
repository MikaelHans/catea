package session

import (
	"context"
	"time"

	"github.com/MikaelHans/catea/session-management/pkg/session"
	"github.com/MikaelHans/catea/user/pkg/structs"
)


func SetSession(token string , memberData structs.Member)(string, error){
	conn, err := connect()

	if err != nil{
		return "", err
	}
	defer conn.Close()

	c := session.NewSessionManagementClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// r, err := c.GetSessionInfo(ctx, &session.SessionID{Sessionid: "asd"})
	data, err := c.SetSession(ctx, &session.SessionData{
		SessionID: &session.SessionID{Sessionid: token}, 
		Email: memberData.Email,
		Firstname: memberData.Firstname,
		Lastname: memberData.Lastname,
		Membersince: memberData.Member_Since,})
	
	return data.String(), err
}

func GetSession(token string)(string, error){
	conn, err := connect()

	if err != nil{
		return "", err
	}
	defer conn.Close()

	c := session.NewSessionManagementClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// r, err := c.GetSessionInfo(ctx, &session.SessionID{Sessionid: "asd"})
	data, err := c.GetSessionInfo(ctx, &session.SessionID{Sessionid: token})
	
	return data.GetData(), err
}