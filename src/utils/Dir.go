package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

func IsDirEmpty() {
	// get current user
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// path to Trash directory `$HOME/.Trash/`
	path := fmt.Sprintf(usr.HomeDir + "/.Trash/")

	// read directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	// get files from directory
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}

	// is Trash directory empty
	if len(files) == 0 {
		fmt.Println("Recycle bin is empty")
	}
}
