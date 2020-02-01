package main

import (
	"fmt"

	"github.com/rrkrish561/relief-call-logger/Caller"
)

func main() {
	fmt.Println("hi")

	newCaller := Caller.Caller{
		CallId: "hi",
	}
	fmt.Println(newCaller)

}
