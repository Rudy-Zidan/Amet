package metis

import (
	"fmt"
	"metis/crawler"
	"net"
	"os"
	"strconv"
)

//Server Configuration
const (
	ConnHost = "localhost"
	ConnPort = "3333"
	ConnType = "tcp"
)

// Start up Metis server
func Start() {
	listener := getListener()
	defer listener.Close()
	startRecieveConnection(listener)
}

func getListener() (conn net.Listener) {
	l, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println("Error Listening:", err.Error())
		os.Exit(1)
	}
	return l
}

func startRecieveConnection(listener net.Listener) {
	fmt.Println("Listening on " + ConnHost + ":" + ConnPort)
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err.Error())
			os.Exit(1)
		}
		fmt.Println(connection.LocalAddr(), "Has been connected")
		go handleRequest(connection)
	}
}

func handleRequest(connection net.Conn) {
	buf := make([]byte, 1024)
	reqLen, err := connection.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(string(buf) + " Total length:" + strconv.Itoa(reqLen))
	connection.Write(crawler.Retrive(string(buf)))
	connection.Close()
}
