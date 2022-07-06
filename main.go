package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/sirupsen/logrus"
)

const (
	CONN_HOST  = "0.0.0.0"
	START_PORT = 4000
	CONN_TYPE  = "tcp"
)

var listenCount int

func main() {
	flag.IntVar(&listenCount, "n", 2, "set listen count")
	flag.Parse()
	logrus.Infof("listen port count: %d", listenCount)
	for i := 0; i < listenCount; i++ {
		go func(iter int) {
			listen(CONN_TYPE, CONN_HOST, strconv.Itoa(START_PORT+iter))
		}(i)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	<-c
}

// Listen for incoming connections.
func listen(netType, host, port string) {
	l, err := net.Listen(netType, host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + host + ":" + port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	for {
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		// Send a response back to person contacting us.
		message := string(buf[:len])
		conn.Write([]byte(fmt.Sprintf("received message from %s, content is: %s\n", conn.RemoteAddr().String(), message)))
		if string(buf[:4]) == "exit" {
			conn.Write([]byte("bye\n"))
			break
		}
	}
	conn.Close()
}
