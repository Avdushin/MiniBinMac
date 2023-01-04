package utils

import (
	"log"
	"os/user"
)

// Get User's home directory (unix)
func GetHomeDir() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("[ERROR] GetHomeDir: %v", err)
	}
	log.Printf("[INFO] Home folder: %s\n", usr.HomeDir)
}
