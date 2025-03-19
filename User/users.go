package user

import (
	"auraluvsu.com/Utils"
	"bufio"
	"fmt"
	"log"
)

type User struct {
	Name    string
	Userid  []byte
	isAdmin bool
}

type ChatMsg struct {
	Message []byte
	User    User
}

const founderKey = "Founder's Key!" // Demo Admin Key

func CreateUser(username, key string) *User {
	UserID := CreateUserID(username) // Creates userID and passes it to the struct
	newUser := &User{                // Create the user and populate the struct
		Name:    username,
		Userid:  UserID,
		isAdmin: false,
	}
	return newUser // Returns the created user
}

func CreateUserID(name string) []byte {
	ID, err := utils.RandBytes(len(name)) // Creates a random byte slice based on the length of the name
	if err != nil {
		log.Fatal("Error processing bytes:", err) // Error handling
	}
	hashedID, err := utils.CreateNewHash(ID) // Creates a bcrypt hash using the byte slice from before
	if err != nil {
		log.Fatal("Failed to hash UserID:", err) // More error handling
	}
	return hashedID // Returns the bcrypt hashed ID
}

func CreatePassword(name string) []byte {
	pswd, err := utils.RandBytes(len(name)) // Creates a random byte slice based on the length of the name
	if err != nil {
		log.Fatal("Error processing bytes:", err) // Error handling
	}
	hashedPSWD, err := utils.CreateNewHash(pswd) // Creates a bcrypt hash using the byte slice from before
	if err != nil {
		log.Fatal("Failed to hash password:", err) // More error handling
	}
	return hashedPSWD // Returns the bcrypt hashed password
}

func (u *User) GetInfo() string {
	userInfo := fmt.Sprintf("Username: %v\nUser ID: %v\n", u.Name, u.Userid)
	return userInfo
}
