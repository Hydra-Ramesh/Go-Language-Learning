package auth

import "fmt"

// We should write Function names starting with Capital letter to make them public
func LoginUserCredentials(username string, password string) {
	fmt.Printf("Logging in user: %s with password: %s\n", username, password)
}
