package repository

import(
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