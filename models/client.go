package models

import (
	"encoding/gob"
	"net"
)

// Client ...
type Client struct {
	network string
	address string
}

// NewClient ...
func NewClient(network string, address string) Client {
	return Client{
		network: network,
		address: address,
	}
}

// Request ...
func (client Client) Request(request Request) (Response, error) {
	conn, err := net.Dial(client.network, client.address)
	if err != nil {
		return Response{}, err
	}

	encoder := gob.NewEncoder(conn)
	encoder.Encode(request)

	var response Response
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&response)
	return response, nil
}
