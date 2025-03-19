package chatter

import (
	"auraluvsu.com/User"
	"fmt"
	"log"
	"net"
	"net/http"

	"auraluvsu.com/Utils"
	"github.com/gorilla/websocket"
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
	msgChan := make(chan string)
	go SendMessage(ws, msgChan)
	go ReceiveMessage(ws, msgChan)
	select {}

}

func SendMessage(conn *websocket.Conn, msgChan chan string) {
	for msg := range msgChan {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Error writing to server:", err)
			return
		}
	}
}

func ReceiveMessage(u user.User, ws *websocket.Conn, msgChan chan string) {
	for {
		_, content, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			close(msgChan)
			return
		}
		log.Printf("%v: %s\n", u.Name, content)
		msgChan <- string(content)
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
