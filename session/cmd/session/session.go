package main

import (
	"context"
	"encoding/json"
	"time"
	"github.com/MikaelHans/catea/session/pkg/structs"
	"github.com/go-redis/redis/v8"
    "github.com/MikaelHans/catea"
)

func ConnectToRedisClient() *redis.Client{
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:8800",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
    
    return client
}

func CreateSession(userID string, sessid string) (*structs.Session, error) {
    
    client := ConnectToRedisClient()
    sessionID := sessid 
    session := &structs.Session{
        ID:     sessionID,
        UserID: userID,
    }

    // Store the session data in Redis
    err := StoreSessionData(client, session)
    if err != nil {
        return nil, err
    }

    return session, nil
}

func StoreSessionData(client *redis.Client, session *structs.Session) error {
    // Serialize the session data (e.g., to JSON)
    ctx := context.Background()
    sessionData, err := json.Marshal(session)
    if err != nil {
        return err
    }

    // Store the session data in Redis with a TTL (time-to-live)
    err = client.Set(ctx, session.ID, sessionData, time.Hour).Err()
    if err != nil {
        return err
    }

    return nil
}
