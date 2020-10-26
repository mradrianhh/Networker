package server

import (
	"Networker/models"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

// Run starts the service.
func Run(network string, service string) {
	go listen(network, service)
}

// Messages holds a record of all the messages passed.
var Messages []models.Message

// Log holds a record of all logs written.
var Log []string

// ErrorLog holds a record of all errors encountered.
var ErrorLog []error

// Conns holds a record of all the connections.
var Conns []net.Conn

func listen(network string, service string) {
	tcpAddr, err := net.ResolveTCPAddr(network, service)
	checkError(err)

	listener, err := net.ListenTCP(network, tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			addError(err)
			continue // if there is an error creating a connection, we'll move on and let the client deal with it. They can retry.
		}

		Conns = append(Conns, conn)

		decoder := gob.NewDecoder(conn)
		var message models.Message
		decoder.Decode(&message) // Read into the "message"-variable from the stream.
		Messages = append(Messages, message)
	}
}

func checkError(err error) {
	if err != nil {
		addError(err)
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func writeLog(text string) {
	Log = append(Log, text)
}

func addError(err error) {
	ErrorLog = append(ErrorLog, err)
}
