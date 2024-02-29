package helper

import (
	"log"
	"os"
)

func Log(fileName, message string) {
	// Open or create a log file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()

	// Set the log output to the file
	log.SetOutput(file)

	// Write log message to the file
	log.Println(message)
}
