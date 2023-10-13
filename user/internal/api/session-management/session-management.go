package api

import (
	"context"
	"time"

	"github.com/MikaelHans/catea/session-management/pkg/session"
	"github.com/MikaelHans/catea/session-management/pkg/structs"
)


func SetSession(structs.Member)(error){
	conn, err := connect()

	if err != nil{
		return err
	}
	defer conn.Close()

	c := session.NewSessionManagementClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// r, err := c.GetSessionInfo(ctx, &session.SessionID{Sessionid: "asd"})
	data, err := c.SetSession(ctx, &session.SessionData{
		SessionID: &session.SessionID{Sessionid:  "asd"}, 
		Email: "mikael.hans30@gmail.com",
		Firstname: "Jonathan",
		Lastname: "Cahyo",
		Membersince: "2022-05-07 17:39:05",})
	
	return err
}