package client

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/mradrianhh/Networker/models"
)

// Request sends a message and returns the response.
func Request(message models.Message, service string, network string) (models.Message, error) {
	if network == "" {
		network = "tcp"
	}

	conn, err := net.Dial(network, service)

	if err != nil {
		return models.NilMessage, err
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
