package macbin

import (
	"log"
	"os"
	"path/filepath"
)

func IsLoaded() {
	log.Printf("Macbin is loaded!")
}

func RemoveGlob(path string) (err error) {
	contents, err := filepath.Glob(path)
	if err != nil {
		return
	}
	for _, item := range contents {
		err = os.RemoveAll(item)
		if err != nil {
			return
		}
		log.Printf("Removed: %s", item)
	}
	return
}
