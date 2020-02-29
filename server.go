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
	if method == "GET" && protocol == "HTTP/1.1" {
		switch request {
		case "/":
			handleRequest(conn,"operations.html","text/html")
			log.Printf("response on: %s", request)
		case "/files/index.html":
			handleRequest(conn,"files/index.html","text/html")
			log.Printf("response on: %s", request)
		case "/files/pngImage.png":
			handleRequest(conn,"files/pngImage.png","image/png")
			log.Printf("response on: %s", request)
		case "/files/jpgImage.jpg":
			handleRequest(conn,"files/jpgImage.jpg","image/png")
			log.Printf("response on: %s", request)
		case "/files/someFile.txt":
			handleRequest(conn,"files/someFile.txt","text/html")
			log.Printf("response on: %s", request)
		case "/files/fizic.pdf":
			handleRequest(conn,"files/fizic.pdf","application/pdf")
			log.Printf("response on: %s", request)
		}
		}
}

func handleRequest(conn net.Conn, fileName,contentType string)  {
		writer := bufio.NewWriter(conn)
		bytes, err := ioutil.ReadFile(fileName)
		_, err= writer.WriteString("HTTP/1.1 200 OK\r\n")
		if err != nil {
			log.Printf("cant  write serv:%v",err)
		}
		_, err = writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(bytes)))
		if err != nil {
			log.Printf("cant write content: %v",err)
		}
		_, err = writer.WriteString("Content-Type:" +contentType+"\r\n")
		if err != nil {
			log.Printf("cant write type: %v",err)
		}
		_, err = writer.WriteString("Connection: Close\r\n")
		if err != nil {
			log.Printf("cant write conn:%v",err)
		}
		_, err = writer.WriteString("\r\n")
		if err != nil {
			log.Printf("cant write: %v",err)
		}
		_, err = writer.Write(bytes)
		if err != nil {
			log.Printf("cant write :%v",err)
		}
		err = writer.Flush()
		if err != nil {
			log.Printf("can't sent response: %v", err)
		}
}