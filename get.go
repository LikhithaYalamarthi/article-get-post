package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {

    resp, err := http.Get("https://inshorts.com/")

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {

        panic(err)
    }

    fmt.Println(string(body))
}