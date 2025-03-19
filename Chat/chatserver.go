package chatter

import (
	"auraluvsu.com/Utils"
	"fmt"
	"log"
	"net"
)

type Chatroom struct {
	ID   []byte
	name string
	port string
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
