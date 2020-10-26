package client

import (
	"Networker/models"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

// MessageServer sends a message to the server and returns the response.
func MessageServer(message models.Message, service string, network string) (models.Message, error) {
	if network == "" {
		network = "tcp"
	}

	conn, err := net.Dial(network, service)

	if err != nil {
		return models.ErrorMessage, err
	}

	encoder := gob.NewEncoder(conn)
	encoder.Encode(message)

	var response models.Message
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&response)
	return response, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
