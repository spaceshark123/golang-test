package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"crypto/sha256"
	"encoding/hex"
)

var wg = sync.WaitGroup{}

var DB map[string]string = make(map[string]string) // username -> hash
var usernameBank []string = []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Heidi", "Ivan", "Judy"}

func main() {
	initializeDB(usernameBank)

	N := 100 // number of queries

	// query the database sequentially and measure time
	start := time.Now()
	for _ = range N {
		wg.Add(1)
		username := usernameBank[rand.Intn(len(usernameBank))] // pick a random username
		queryDB(username)
	}
	end := time.Now()
	fmt.Println("Sequential query time: ", end.Sub(start))

	// query the database concurrently and measure time
	start = time.Now()
	for _ = range N {
		wg.Add(1)
		username := usernameBank[rand.Intn(len(usernameBank))] // pick a random username
		go queryDB(username)
	}
	wg.Wait()
	end = time.Now()
	fmt.Println("Concurrent query time: ", end.Sub(start))
}

func queryDB(username string) {
	defer wg.Done()
	delay := rand.Intn(100) // sleep for a random amount of time to simulate a query
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("username: %s -> %s\n", username, DB[username])
}

func initializeDB(usernameBank []string) {
	for _, username := range usernameBank {
		DB[username] = Hash(username)
	}
}

// hash a username using SHA-256
func Hash(username string) string {
	hash := sha256.New()
	hash.Write([]byte(username))
	return hex.EncodeToString(hash.Sum(nil))
}