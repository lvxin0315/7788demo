package main

/*
#include <a.c>
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	a.i = 1
	a.f = 1.1
	fmt.Println(a.i)
	fmt.Println(a.f)
}
