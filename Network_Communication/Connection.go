import (
    "net"
    "bufio"
)

func startServer(port string) {
    ln, _ := net.Listen("tcp", ":"+port)
    defer ln.Close()

    for {
        conn, _ := ln.Accept()
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println("Received:", message)
    conn.Close()
}

func connectToPeer(address string) {
    conn, _ := net.Dial("tcp", address)
    fmt.Fprintf(conn, "Hello from fadaka\n")
    conn.Close()
}
