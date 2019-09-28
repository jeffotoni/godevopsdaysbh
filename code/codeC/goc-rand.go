package main

import (
	"fmt"
)

/*
#include <stdlib.h>
*/
import "C"

func Random() int64 {
	return int64(C.rand())
}

func main() {
	fmt.Println(Random())
}
