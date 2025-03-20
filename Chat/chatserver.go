package chatter

import (
	"auraluvsu.com/User"
	"auraluvsu.com/Utils"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

// This struct is created for ease of storing information about each created room
type Chatroom struct {
	ID   []byte
	name string
	port string
}

// Websocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer ws.Close()
	log.Println("Websocket connection established!")
	var newUsername string
	var userKey string
	fmt.Println("Enter new Username:")
	fmt.Scan(&newUsername)
	fmt.Println("Enter optional key:")
	fmt.Scan(&userKey)
	newUser := user.CreateUser(newUsername, userKey)
	msgChan := make(chan string)
	go SendMessage(ws, newUser.Name, msgChan)
	go ReceiveMessage(ws, msgChan)
	select {}

}

func SendMessage(conn *websocket.Conn, username string, msgChan chan string) {
	for msg := range msgChan {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Error writing to server:", err)
			return
		}
	}
}

func ReceiveMessage(ws *websocket.Conn, msgChan chan string) {
	_, nameTxt, err := ws.ReadMessage()
	if err != nil {
		log.Println("Error getting username:", err)
		return
	}
	username := string(nameTxt)
	log.Printf("User '%s' is connected", username)

	for {
		_, content, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			close(msgChan)
			return
		}
		log.Println(fmt.Sprintf("|| [%v]: %s\n", username, content))
		msgChan <- fmt.Sprintf("|| [%v]: %s\n", username, string(content))
	}
}

func KillServer() {
	ok, conn := SecuredConn()
	if ok {
		log.Println("Server Terminated...")
		conn.Close()
	} else {
		log.Println("Server Already Terminated...")
		return
	}
}

func SecuredConn() (bool, net.Conn) {
	conn, err := SetConnection(8080)
	if err != nil {
		log.Fatal(err)
	}
	if conn != nil {
		return true, conn
	} else {
		return false, nil
	}
}
func CreateNewRoom(newName string) Chatroom {
	idBytes, err := utils.RandBytes(8)
	if err != nil {
		log.Fatalf("Error getting custom ID: %v", err)
	}

	newRoom := Chatroom{
		ID:   idBytes,
		name: newName,
	}
	return newRoom
}

func SetConnection(port int16) (net.Conn, error) {
	host := "192.168.0.100"
	address := fmt.Sprintf("%v:%v", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to server...")
	}
	fmt.Printf("Connected to server: %v\n", port)
	return conn, nil
}
