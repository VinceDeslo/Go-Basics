package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Employee data format
type Employee struct {
	Name       string `json:"Name"`
	Department string `json:"Dep"`
	Salary     int    `json:"Salary"`
}

func main() {

	fmt.Println("Starting up Redis client.")
	var ctx = context.Background()

	// Instantiate the redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	// Encode an entry into json
	json, err := json.Marshal(Employee{Name: "Richard", Department: "IT", Salary: 50000})
	if err != nil {
		fmt.Println(err)
	}

	// Store the data into redis
	err = rdb.Set(ctx, "id1", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Retrieve the data
	val, err := rdb.Get(ctx, "id1").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
