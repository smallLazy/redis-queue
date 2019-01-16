package main

import (
	"log"
	_ "redis_queue/routers"

	"github.com/go-redis/redis"
)

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	pong, err := client.Ping().Result()
	log.Println(pong, err)
	return client
}
func main() {
	client := GetRedisClient()
	err := client.Set("name", "lazy", 0).Err()
	if err != nil {
		log.Println(err)
	}
	value, err := client.Get("name").Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("name=:", value)
}
