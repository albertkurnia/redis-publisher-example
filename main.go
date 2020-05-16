package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

var (
	channel = "EXAMPLE"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}

	pubsub := client.Subscribe(channel)
	_, err = pubsub.Receive()
	if err != nil {
		fmt.Println(err)
	}

	// PUBLISH
	counter := 0
	for counter < 10 {
		s := rand.NewSource(time.Now().Unix())
		message := s.Int63()

		if counter == 9 {
			message = -1
		}
		_ = client.Publish(channel, message)
		fmt.Println("Publish: ", message)
		counter++
		time.Sleep(time.Second)
	}

	_ = pubsub.Close()
	fmt.Println("Pubsub closed. Channel closed.")
}
