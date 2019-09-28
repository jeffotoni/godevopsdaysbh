package main

/*
#include "randnew.h"
*/
import "C"

func main() {

	C.Print(C.randnew())
}
