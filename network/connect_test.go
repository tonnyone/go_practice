package network

import (
	"log"
	"net"
	"testing"
	"time"
)

// start a tcp server
func setupTCP() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	log.Println("Listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Println("a client connected")
}

func TestTcpConnectCorrect(t *testing.T) {
	go setupTCP()
	time.Sleep(5 * time.Second)
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		t.Fatal("could not connect to server: ", err)
	}
	defer conn.Close()
}

func TestTcpConnectMistake(t *testing.T) {
	go setupTCP()
	time.Sleep(5 * time.Second)
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		t.Fatal("could not connect to server: ", err)
	}
	defer conn.Close()
}
