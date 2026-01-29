package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"github.com/redis/go-redis/v9" // Import the Redis library
)

var ctx = context.Background()

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. Connect to Redis (using the Kubernetes Service Name "redis-service")
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis-service:6379", 
		Password: "", // no password set
		DB: 0,  // use default DB
	})

	// 2. Increment the "hits" counter
	val, err := rdb.Incr(ctx, "hits").Result()
	if err != nil {
		fmt.Fprintf(w, "Redis Error: %v", err)
		return
	}

	// 3. Get Pod Name
	hostname, _ := os.Hostname()
	
	// 4. Print the result
	fmt.Fprintf(w, "Hello Network Engineer! I am running on Pod: %s\n", hostname)
	fmt.Fprintf(w, "Total Visitor Count: %d\n", val)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
