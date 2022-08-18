package main

import (
	"log"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "[my-logging] ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	logger.Println("### START ###")

	logger.Println("--- END ---")
}
