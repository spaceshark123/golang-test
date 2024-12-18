package main

import (
	"github.com/spaceshark123/golang-test/src/database"
	"bufio"
	"os"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

// accounts mock database storing username -> password hash
var accounts database.MockDB[string, string] = database.MockDB[string, string]{}

func main() {
	accounts.Init() // initialize the database
	defer accounts.Close() // close the database when done

	// add some accounts
	Register("admin", "password123");
	Register("johnsmith", "123456");
	Register("testuser", "abc123");
	Register("user123", "str0ngPa5sw0rd123");
	Register("spaceshark123", "golangIsAwesome");

	// provide CLI for user to login
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter username: ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}
		// login
		if UserExists(input) {
			fmt.Println("Enter password: ")
			scanner.Scan()
			password := scanner.Text()
			if Login(input, password) {
				fmt.Println("Login successful!")
			} else {
				fmt.Println("Login failed!")
			}
		} else {
			fmt.Println("User does not exist!")
		}
		// prompt for username again
		fmt.Println("Enter username: ")
	}
}

func UserExists(username string) bool {
	_, err := accounts.Get(username)
	return err == nil
}

func Register(username string, password string) {
	// store the password hash in the database
	accounts.Set(username, Hash(password))
}

func Login(username string, password string) bool {
	// get the password hash from the database
	hash, err := accounts.Get(username)
	if err != nil {
		return false
	}

	// check if the password h ashmatches the stored hash
	return Hash(password) == hash
}

// hash a username using SHA-256
func Hash(username string) string {
	hash := sha256.New()
	hash.Write([]byte(username))
	return hex.EncodeToString(hash.Sum(nil))
}