package main

import (
	"bufio"
	"os"
)

// write bufio to optimization
var writer *bufio.Writer

func Println(text string) {
	writer = bufio.NewWriter(os.Stdout)
	writer.WriteString("\n")
	writer.WriteString(text)
	writer.Flush()
}

func main() {
	Println("Ola Workshop2.0 ...")
}
