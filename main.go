package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	// Start the server
	fmt.Println("Starting the server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start TCP Server: %v", err)
	}
	defer listener.Close()
	log.Println("TCP Server started on port 8080")
	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept incoming connection: %v", err)
			continue
		}
		go handleConnection(conn)
		//go handleConnection2(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read HTTP request from the connection
	request, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Printf("Failed to read HTTP request: %v", err)
		return
	}

	// Extract necessary information from the HTTP request
	method := request.Method
	url := request.URL.String()
	headers := request.Header
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v", err)
		return
	}

	// Create a new HTTP request based on the extracted information
	newRequest, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		log.Printf("Failed to create new HTTP request: %v", err)
		return
	}
	newRequest.Header = headers

	// Forward the new HTTP request to the target server
	response, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		log.Printf("Failed to forward request to target server: %v", err)
		return
	}
	defer response.Body.Close()

	// Handle the response from the target server
	// You can forward the response back to the client or process it as needed
}

func handleConnection2(conn net.Conn) {

}
