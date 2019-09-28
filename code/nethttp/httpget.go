// Go Api server
// @jeffotoni

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {

    response, err := http.Get("http://golang.org/")
    if err != nil {
        fmt.Printf("\n%s", err)
        os.Exit(0)
    }

    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("\n%s", err)
        os.Exit(0)
    }

    fmt.Printf("%s\n", string(contents))
}
