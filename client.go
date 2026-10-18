package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Connection error: ", err)
		return
	}

	defer conn.Close()

	fmt.Println("Connected to server!")

	reader := bufio.NewReader(conn)

	resp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
    	return
	}

	fmt.Printf("Server response: %q\n", resp)

	if resp != "OK\n" {
		fmt.Println("Invalid server response!")
    	return
	}

	fmt.Println("Server response is correct!")
}