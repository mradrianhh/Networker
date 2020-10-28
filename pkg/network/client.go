package network

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/mradrianhh/Networker/pkg/network/models"
)

// Client ...
type Client struct {
	network  string
	address  string
	errors   []error
	handlers map[models.RequestCode]Handler
}

// NewClient ...
func NewClient(network string, address string) Client {
	return Client{
		network:  network,
		address:  address,
		errors:   []error{},
		handlers: make(map[models.RequestCode]Handler),
	}
}

// AddHandler ...
func (client *Client) AddHandler(requestCode models.RequestCode, handler Handler) {
	client.handlers[requestCode] = handler
}

// Request ...
func (client Client) Request(request models.Request) (models.Response, error) {
	conn, err := net.Dial(client.network, client.address)
	if err != nil {
		return models.Response{}, err
	}

	encoder := gob.NewEncoder(conn)
	encoder.Encode(request)

	var response models.Response
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&response)
	return response, nil
}

// Respond ...
func (client Client) Respond(response models.Response, conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(response)
}

// Listen ...
func (client *Client) Listen() {
	tcpAddress, err := net.ResolveTCPAddr(client.network, client.address)
	client.checkError(err)

	listener, err := net.ListenTCP(client.network, tcpAddress)
	client.checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			client.errors = append(client.errors, err)
			continue
		}

		fmt.Println(conn.LocalAddr())

		decoder := gob.NewDecoder(conn)
		var request models.Request
		decoder.Decode(&request)
		fmt.Println(request.RequestCode)
		if err := client.handlers[request.RequestCode](request, conn); err != nil {
			client.Respond(models.NewResponse(models.ERROR), conn)
			client.errors = append(client.errors, err)
			continue
		} else {
			client.Respond(models.NewResponse(models.CONFIRMATION), conn)
		}
	}
}

func (client *Client) checkError(err error) {
	if err != nil {
		client.errors = append(client.errors, err)
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
