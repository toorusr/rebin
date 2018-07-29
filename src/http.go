package main

import (
	"io"
	"net/http"
    "fmt"
    "time"
    "os"
    "io/ioutil"
)
const (
    postsDir = "./posts.ign/"
)

func rootView(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
    io.WriteString(w, "Accesstime: " + t.Format("2006-01-02 15:04:05") + "\n\n\n")
    if len(r.URL.Path[1:]) > 0 {
        io.WriteString(w, "\nYou send more than nothing: " + r.URL.Path[1:])
        data := dataHandler(r.URL.Path[1:])
        if data != "nil" {
            io.WriteString(w, "\nAnd your extra data is a post: \n\n" + data)
        } else {
            io.WriteString(w, "\nBut your extra data is useless..")
        }
    } else {
        io.WriteString(w, "Usage:\n\techo test | nc localhost 1337\n\n")
    }
}

func statsView(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Analytics page</h1>")
    files, err := ioutil.ReadDir(postsDir)
    if err != nil {
        fmt.Print(err)
    }

    for _, f := range files {
        io.WriteString(w, f.Name() + "\n")
    }
}

func breadView(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Yummy;")
}

func read(id string) string {
    b, err := ioutil.ReadFile(postsDir + id) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    return string(b)
}

func dataHandler(id string) string {
    if _, err := os.Stat(postsDir + id); err == nil {
        return read(id)
    } else {
        return "nil"
    }
}

func main() {
	http.HandleFunc("/", rootView)
    http.HandleFunc("/bread", breadView)
    http.HandleFunc("/stats", statsView)
    fmt.Println("Listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
