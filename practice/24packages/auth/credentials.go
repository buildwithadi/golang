package auth

import "fmt"

// lowercased function : is private
// Capitalize function : is public
func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, password)
}