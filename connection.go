package main

import (
    "crypto/tls"
    "fmt"
    "net"
    "os"
)

func startSecurePeerConnection(address string) {
    cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        fmt.Println("Error loading certificates:", err)
        os.Exit(1)
    }

    config := &tls.Config{Certificates: []tls.Certificate{cert}}
    ln, err := tls.Listen("tcp", address, config)
    if err != nil {
        fmt.Println("Error starting server:", err)
        os.Exit(1)
    }
    defer ln.Close()

    fmt.Println("Secure peer connection started on", address)

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handlePeerConnection(conn)
    }
}
