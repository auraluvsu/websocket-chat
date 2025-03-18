package user

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"auraluvsu.com/Admin"
	"auraluvsu.com/Chat"
	"auraluvsu.com/Utils"
)

type User struct {
	Name    string
	Userid  []byte
	Room    chatter.Chatroom
	isAdmin bool
}

type ChatMsg struct {
	Message []byte
	User    User
}

const founderKey = "Founder's Key!" // Demo Admin Key

func CreateUser(username, key string) *User {
	UserID := CreateUserID(username)          // Creates userID and passes it to the struct
	newAdmin := &admin.Admin{Key: founderKey} // Gonna amend this later to check if the inputted admin key is valid
	newUser := &User{                         // Create the user and populate the struct
		Name:    username,
		Userid:  UserID,
		isAdmin: false,
	}
	if newAdmin.Key == founderKey { // Check for valid admin key
		newUser.isAdmin = true
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

func ChooseRoom(room string) (net.Conn, error) {
	conn, err := net.Dial("tcp", room)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to server: %v", err)
	}
	return conn, nil
}

func SendMessage(conn net.Conn, msgChan chan string) {
	for msg := range msgChan {
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing to server:", err)
			return
		}
	}
}

func ReceiveMessage(conn net.Conn, msgChan chan string) {
	reader := bufio.NewReader(conn)
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected")
			close(msgChan)
			return
		}
		msgChan <- content
	}
}

func (u *User) GetInfo() string {
	userInfo := fmt.Sprintf("Username: %v\nUser Room: %v\n", u.Name, u.Room)
	return userInfo
}
