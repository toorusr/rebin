package main

import (
    "fmt"
    "net"
    "os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "1337"
    CONN_TYPE = "tcp"
)

func main() {
    // Listening for incomming
    lis, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if err != nil {
        fmt.Println("Error listening: ", err.Error())
        os.Exit(1)
    }
    defer lis.Close() // Close app when listener closes
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        conn, err := lis.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    buf := make([]byte, 4096)
    _, err := conn.Read(buf)
    fmt.Println("Rec:\n", string(buf))
    if err != nil {
        fmt.Println("Error reading: ", err.Error())
    }
    conn.Write([]byte("Received message: " + string(buf)))
    conn.Close()
}
