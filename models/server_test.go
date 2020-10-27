package models_test

import (
	"net"
	"testing"

	networker "github.com/mradrianhh/Networker"
	"github.com/mradrianhh/Networker/models"
)

func TestNewServer(t *testing.T) {
	got := models.NewServer("tcp", "0.0.0.0:0")
	expected := models.Server{
		network:  "tcp",
		address:  "0.0.0.0:0",
		conns:    []net.Conn{},
		errors:   []error{},
		messages: []Message{},
		handlers: make(map[networker.RequestCode]handler),
	}

	if got != expected {
		t.Errorf("NewServer did not return the expected result.")
	}
}
