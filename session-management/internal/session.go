package session

import (
	"context"
	"encoding/json"
	"time"

	repo "github.com/MikaelHans/catea/session-management/internal/repository"
	"github.com/MikaelHans/catea/session-management/pkg/structs"
	pb "github.com/MikaelHans/catea/session-management/api"
)

type Server struct {
	pb.UnimplementedSessionManagementServer
}

func (s *Server) GetSessionInfo(ctx context.Context, sessionID *pb.SessionID) (*pb.String, error) {
	client := repo.ConnectToRedisClient()
	data, err := client.Get(ctx, sessionID.GetSessionid()).Result()
    var tmp pb.String   
	if err != nil {
        tmp.Data = data
		return &tmp, err
	}
	return &tmp, err
}

func (s *Server) SetSession(ctx context.Context, sessionData *pb.SessionData)(*pb.Empty ,error){
	client := repo.ConnectToRedisClient()
	data, err := json.Marshal(sessionData)
	if err != nil {
		return &pb.Empty{},err
	}
	err = client.Set(ctx, sessionData.SessionID.Sessionid, data, time.Hour).Err()
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}

func StoreSessionDataToRedis(memberdata structs.Member, token string, ctx context.Context) error {
	client := repo.ConnectToRedisClient()
	sessionData, err := json.Marshal(memberdata)
	if err != nil {
		return err
	}
	err = client.Set(ctx, token, sessionData, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}
