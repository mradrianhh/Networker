package client

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/mradrianhh/Networker/models"
)

// Request sends a request and returns the response.
func Request(request models.Request, service string, network string) (models.Response, error) {
	if network == "" {
		network = "tcp"
	}

	conn, err := net.Dial(network, service)

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

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
