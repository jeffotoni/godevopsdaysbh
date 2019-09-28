// Go in action
// @jeffotoni
// 2019-08-10

package main

import (
  "fmt"
  "os"
  "plugin"
)

var (
  mod = "../plugin/tolower.so"
)

type TolowerID interface {
  Tolower(name string) string
}

func main() {
  // load module
  // 1. open the so file to load the symbols
  plug, err := plugin.Open(mod)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // 2. look up a symbol (an exported function or variable)
  // in this case, variable Greeter
  tlower, err := plug.Lookup("Tolower")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // 3. Assert that loaded symbol is of a desired type
  // in this case interface type Greeter (defined above)
  var tl TolowerID
  tl, ok := tlower.(TolowerID)
  if !ok {
    fmt.Println("unexpected type from module symbol")
    os.Exit(1)
  }

  // 4. use the module
  fmt.Println(tl.Tolower("@JEFFOTONI/LAMBDAMAN"))
}
