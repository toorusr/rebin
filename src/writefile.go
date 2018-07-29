package main

import (
    "io/ioutil"
)

const (
    postsDir = "./posts.ign/"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}



func main() {
    write("test", "Test-String")
}
