package main

import (
	"context"
	"encoding/json"
	"time"
	"github.com/go-redis/redis/v8"
)

func ConnectToRedisClient() *redis.Client{
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:8800",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })

    return client
}

func createSession(userID string, sessid string) (*Session, error) {

    client := ConnectToRedisClient()
    sessionID := sessid // You need to implement this function
    session := &Session{
        ID:     sessionID,
        UserID: userID,
    }

    // Store the session data in Redis
    err := storeSessionData(client, session)
    if err != nil {
        return nil, err
    }

    return session, nil
}

func storeSessionData(client *redis.Client, session *Session) error {
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
