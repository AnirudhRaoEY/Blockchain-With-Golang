package main

import (
	"fmt"
	"log"
)
//Logging Function
func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	a := "Test"
	fmt.Println(a)
	log.Println("Test 2")
}
