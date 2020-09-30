package main

// #include <unistd.h>
import "C"
import (
	"fmt"
)

func main() {
	go func() {
		C.sleep(1000000000)
		fmt.Println("I slept")
	}()
	fmt.Println("Hello world")
}