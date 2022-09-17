package main

import (
	"learning_gorm/crud"
	"log"
)

func main() {
	err := crud.SearchAll()
	if err != nil {
		log.Fatalln(err)
	}
}
