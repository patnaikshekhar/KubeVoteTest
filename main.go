package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func main() {

	client := connectToCache()

	http.HandleFunc("/dogs", createHandleFunc(client, "dogs"))
	http.HandleFunc("/cats", createHandleFunc(client, "cats"))

	log.Print("Server Started")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func connectToCache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "cache:6379",
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Print("Connected to cache")

	return client
}

func createHandleFunc(client *redis.Client, key string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		result := Response{}

		value, err := client.Incr(key).Result()
		if err != nil {
			result.Error = err.Error()
		} else {
			result.Count = int(value)
		}

		jsonResponse, err := json.Marshal(result)
		if err != nil {
			result.Error = err.Error()
		}

		w.Write([]byte(jsonResponse))
	}
}

type Response struct {
	Count int
	Error string
}
