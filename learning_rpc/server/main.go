package main

import "log"

func main() {
	_, err := StartRPCServer()
	if err != nil {
		log.Fatal(err)
	}
}
