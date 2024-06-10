package main

// #include <stdio.h>
//
// void print_hello() {
//     printf("Hello from C!\n");
// }
import "C"
import "fmt"

func main() {
	fmt.Println("Calling C function...")
	C.print_hello()
	fmt.Println("C function called successfully!")
}
