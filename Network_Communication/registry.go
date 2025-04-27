package main

import (
	"fmt"
	"net"
	"sync"
)

var peers = make(map[string]string)
var mutex = &sync.Mutex{}

// Register a peer
func registerPeer(name, address string) {
	mutex.Lock()
	defer mutex.Unlock()
	peers[name] = address
	fmt.Printf("Registered peer: %s at %s\n", name, address)
}

// Get all registered peers
func getPeers() map[string]string {
	mutex.Lock()
	defer mutex.Unlock()
	return peers
}

// Server to handle peer registrations and lookups
func startRegistryServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting registry server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Registry server started on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleRegistryConnection(conn)
	}
}

// Handle peer registration and queries
func handleRegistryConnection(conn net.Conn) {
	defer conn.Close()

	// Receive message from peer
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	message := string(buffer[:n])
	if message == "register" {
		// Register peer
		conn.Write([]byte("Please send your name:"))
		n, _ = conn.Read(buffer)
		name := string(buffer[:n])
		conn.Write([]byte("Please send your address:"))
		n, _ = conn.Read(buffer)
		address := string(buffer[:n])

		registerPeer(name, address)
		conn.Write([]byte("Registration successful\n"))
	} else if message == "getPeers" {
		// Send list of peers
		peers := getPeers()
		for name, address := range peers {
			conn.Write([]byte(fmt.Sprintf("%s - %s\n", name, address)))
		}
	} else {
		conn.Write([]byte("Unknown command"))
	}
}

func main() {
	startRegistryServer("9090")
}
