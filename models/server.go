package models

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	networker "github.com/mradrianhh/Networker"
)

/*
Brief description:

The server will be an object which can listen for requests, and respond accordingly.
The server has a network type(tcp/udp) and an ip-address.
The server has a list of it's connections.
The server has a list of the errors that has occured.
The server has a list of the messages that has been passed.
The server also needs a way to respond(i.e callbacks/handler-methods). How can I implement the callback mechanism?
*/

type handler func(Message, net.Conn) error

// Server ...
type Server struct {
	network  string
	address  string
	conns    []net.Conn
	errors   []error
	messages []Message
	handlers map[networker.RequestCode]handler
}

// NewServer returns a new server-object. The lists are empty and must be initialized manually.
func NewServer(network string, address string) Server {
	return Server{
		network:  network,
		address:  address,
		conns:    []net.Conn{},
		errors:   []error{},
		messages: []Message{},
		handlers: make(map[networker.RequestCode]handler),
	}
}

// Listen listens for connections, reads the request, and responds with the appropriate handler-method.
func (server *Server) Listen() {
	tcpAddress, err := net.ResolveTCPAddr(server.network, server.address)
	server.checkError(err)

	listener, err := net.ListenTCP(server.network, tcpAddress)
	server.checkError(err)

	for {
		fmt.Println("Listening")
		conn, err := listener.Accept()
		if err != nil {
			server.errors = append(server.errors, err)
			continue
		}

		server.conns = append(server.conns, conn)

		decoder := gob.NewDecoder(conn)
		var request Request
		decoder.Decode(&request)
		fmt.Println(request.requestCode)
		server.messages = append(server.messages, request)
		if err := server.handlers[request.requestCode](request, conn); err != nil {
			server.errors = append(server.errors, err)
			continue
		}
	}
}

func (server *Server) checkError(err error) {
	if err != nil {
		server.errors = append(server.errors, err)
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
