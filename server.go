package main

import (
	"fmt"
	"net"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer list.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
    		return
		}

		fmt.Println("Client connected!")

		_, err = conn.Write([]byte("OK\n"))
		if err != nil {
			fmt.Println("Error sending response:", err)
		}

		conn.Close()
	}
}