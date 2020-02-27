package main

import (
	"bytes"
	"io/ioutil"
	"net"
	"strings"
	"testing"
	"time"
)

func Test_forServer(t *testing.T) {

	go func() {
		err := operations("localhost:9999")
		if err != nil {
			t.Fatalf("can't start server: %v", err)
		}
	}()

	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		t.Fatalf("can't connect to server: %v", err)
	}
	defer conn.Close()
	var buffer bytes.Buffer

	buffer.Write([]byte("GET / HTTP/1.1\r\nHost: localhost\r\n\r\n"))
	buffer.WriteTo(conn)
	byte, err := ioutil.ReadAll(conn)
	if err != nil {
		t.Fatalf("can't read answer from server: %v", err)
	}
	response := string(byte)
	if !strings.Contains(response, "HTTP/1.1 200 Ok") {
		t.Fatalf("without answer: %s", response)
	}
}