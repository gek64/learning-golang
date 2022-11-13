package main

import (
	"fmt"
	"log"
)

func main() {
	reply, err := StartRPCClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(reply))
}
