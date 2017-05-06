package main

import (
    "net/url"
    "fmt"
)

func main() {
    value, err := url.ParseQuery("x=1&y=2&z=3")
    if err != nil {
        panic(err)
    }
    fmt.Println(value)
}
