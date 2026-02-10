package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client // Define globally to reuse connections
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. Increment the "hits" counter in Redis
	val, err := rdb.Incr(ctx, "hits").Result()
	if err != nil {
		fmt.Fprintf(w, "Redis Error: %v", err)
		return
	}

	// 2. Get Pod Name
	hostname, _ := os.Hostname()
	
	// 3. Print the result
	fmt.Fprintf(w, "Hello DevOps Engineer! Welcome to Jaimin's Fully Automated CI/CD Pipeline!: %s\n", hostname)
	fmt.Fprintf(w, "Total Visitor Count (Persistent): %d\n", val)
}

func main() {
	// 4. Initialize Redis connection once at startup
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "redis-service:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", 
		DB:       0,  
	})

	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
