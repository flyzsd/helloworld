package main

import (
	"fmt"
	"github.com/flyzsd/calculator"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

/**
This is an comment
*/
func main() {
	fmt.Println("Hello World!")
	log.Println("Hello World!")
	sum := calculator.Sum(1, 2)
	fmt.Printf("sum = %d\n", sum)

	TestEmployee()
}
