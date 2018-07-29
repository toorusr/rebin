package main

import (
    "fmt"
    "net"
    "os"
    "strings"
    "bytes"
    "github.com/speps/go-hashids"
    "time"
    "io/ioutil"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "1337"
    CONN_TYPE = "tcp"
    postsDir = "./posts.ign/"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Listening for incomming
    lis, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
    if err != nil {
        fmt.Println("Error listening: ", err.Error())
        os.Exit(1)
    }
    defer lis.Close() // Close app when listener closes
    fmt.Println("Listening on tcp://" + CONN_HOST + ":" + CONN_PORT)
    for {
        conn, err := lis.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        go handleRequest(conn)
    }
}

func hdHandler() string {
    t := time.Now()
    dbid := t.String()
    hd := hashids.NewData()
    hd.Salt = dbid
    h, _ := hashids.NewWithData(hd)
    id, _ := h.Encode([]int{13, 37})
    // _, _ := h.DecodeWithError(id)
    id = strings.ToLower(id)
    return id
}

func read(id string) string {
    b, err := ioutil.ReadFile(postsDir + id) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    return string(b)
}

func write(id string, data string) {
    f, e := os.Create(postsDir + id)
    f.Close()
    check(e)
    err := ioutil.WriteFile(postsDir + id, bytes.Trim([]byte(data), "\x00"), 0644)
    check(err)
    return
}

func dataHandler(data string) string {
    id := hdHandler()
    write(id, data)
    return id
}

func bufferHandler() {
    return
}

func handleRequest(conn net.Conn) {
    buf := make([]byte, 4096)
    _, err := conn.Read(buf)
    fmt.Println("Rec:\n", string(buf))
    if err != nil {
        fmt.Println("Error reading: ", err.Error())
    }
    dh := dataHandler(string(buf))
    conn.Write([]byte("\nGenerated id: " + dh + "\n"))
    conn.Close()
}
