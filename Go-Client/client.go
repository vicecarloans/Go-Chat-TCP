package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')

		if err == io.EOF {
			log.Fatal("Server hung up connection")
		}
		if err != nil {
			log.Fatalf("Unable to read message: %s", err)
		}
		fmt.Print(msg)
	}
}
func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter your name: ")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	nameInput = nameInput[:len(nameInput)-1]

	fmt.Println("********** MESSAGES **********")
	go onMessage(connection)

	connection.Write([]byte(fmt.Sprintf("%s has joined\n", nameInput)))

	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s: %s \n", nameInput, msg[:len(msg)-1])
		connection.Write([]byte(msg))
	}
	connection.Close()
}
