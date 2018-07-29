package main

import (
    "fmt"
    "io/ioutil"
)

const (
    postsDir = "./posts.ign/"
)

func read(id string) string {
    b, err := ioutil.ReadFile(postsDir + id) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    return string(b)
}

func main() {
    fmt.Println(read("test")) // print the content as a 'string'
}
