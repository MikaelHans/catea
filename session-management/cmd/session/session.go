package session

import (
	"context"
	"encoding/json"
	"time"

	// sessionstructs "github.com/MikaelHans/catea/session/pkg/structs"
	// "github.com/MikaelHans/catea/login-signup/pkg/structs"
	"github.com/MikaelHans/catea/session-management/pkg/structs"
	"github.com/go-redis/redis/v8"
)

func connectToRedisClient() *redis.Client{
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:8800",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
    return client
}

// func CreateSession(memberData loginsignupstructs.Member, token string, ctx context.Context) (error) {
//     client := connectToRedisClient()
//     // Store the session data in Redis
//     err := StoreSessionData(client, memberData, token, ctx)
//     if err != nil {
//         return err
//     }
//     return nil
// }

func StoreSessionData(memberdata structs.Member, token string, ctx context.Context) error {
    client := connectToRedisClient()
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
