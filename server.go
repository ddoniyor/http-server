package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	host := "0.0.0.0"
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9999"
	}
	err := operations(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}
}

func operations(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("can't listen %s: %w", addr, err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Can't close conn: %v", err)
		}
	}()
	for {
		conn, err := listener.Accept()
		log.Print("accept connection")
		if err != nil {
			log.Printf("can't accept: %v", err)
			continue
		}
		log.Print("handle connection")
		handleConn(conn)
	}
}


func handleConn(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Can't close conn: %v", err)
		}
	}()
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	log.Print(requestLine)
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) != 3 {
		return
	}

	method, request, protocol := parts[0], parts[1], parts[2]
	if method == "GET" && request == "/" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("operations.html")
		writer.WriteString("HTTP/1.1 200 Ok\r\n")
		writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		writer.WriteString("Content-Type: text/html\r\n")
		writer.WriteString("Connection: close\r\n")
		writer.WriteString("\r\n")
		writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)

	}

	if method == "GET" && request == "/files/index.html" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("files/index.html")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: text/html\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/files/pngImage.png" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("files/pngImage.png")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: image/png\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/files/jpgImage.jpg" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("files/jpgImage.jpg")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: image/png\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/files/someFile.txt" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("files/someFile.txt")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: text/html\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}

	if method == "GET" && request == "/files/fizic.pdf" && protocol == "HTTP/1.1" {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile("files/fizic.pdf")
		_, _ = writer.WriteString("HTTP/1.1 200 OK\r\n")
		_, _ = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		_, _ = writer.WriteString("Content-Type: application/pdf\r\n")
		_, _ = writer.WriteString("Connection: Close\r\n")
		_, _ = writer.WriteString("\r\n")
		_, _ = writer.Write(bytes)
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
		log.Printf("response on: %s", request)
	}
}