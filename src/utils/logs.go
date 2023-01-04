package utils

import (
	"log"
	"os"
)

// logging module
func LogsInit(f string) {
	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Println("[SUCCESS] MiniBinMac has been started...")
	log.Printf("[INFO] Logs are written to file %s", f)
}
