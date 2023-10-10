package session

import (
	"context"
	"encoding/json"
	"time"

	repo "github.com/MikaelHans/catea/session-management/internal/repository"
	"github.com/MikaelHans/catea/session-management/pkg/structs"
	"github.com/MikaelHans/catea/session/cmd/session"
)

type Server struct{
    session.UnimplementedSessionManagementServiceServer
}
func (s *Server)GetSessionInfo(ctx context.Context, sessionID *SessionID) (string, error){
    client := repo.ConnectToRedisClient()
    memberData, err := client.Get(ctx, sessionID.GetSessionid()).Result()
    if err != nil {
        return memberData, err
    }
    return memberData, err
}

func storeSessionDataToRedis(memberdata structs.Member, token string, ctx context.Context) error {
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
