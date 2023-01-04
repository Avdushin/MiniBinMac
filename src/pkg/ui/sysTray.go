package ui

import (
	"MiniBinMac/src/pkg/icons"
	"MiniBinMac/src/pkg/macbin"
	"MiniBinMac/src/pkg/ui/layout"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"time"

	"github.com/getlantern/systray"
)

var sleepTime = 1
var counter = 0

// Render system tray app
func RenderSysTray() {
	log.Println("sysTray module has been loaded...")

	// Menu
	mClearBin := systray.AddMenuItem("Очистить", "Clear the recicle bin")
	systray.AddSeparator()
	go layout.SetupMenu()

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// OnClick Menu -> Clear
	go func() {
		for {
			<-mClearBin.ClickedCh
			log.Println("Recicle bin has been cleared by click")
			macbin.RemoveGlob(usr.HomeDir + "/.Trash/*")
			go IsDirEmpty()
			counter = 0
		}
	}()
	// Backgroud checking is Dir empty
	go func() {
		for {
			go IsDirEmpty()
			time.Sleep(time.Duration(sleepTime) * time.Second)
			counter++

			if counter > 15 {
				sleepTime = 5
			} else if counter < 15 {
				sleepTime = 1
			}
		}
	}()
}

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
	// is Trash directory empty srt `Empty icon` else set `Full cart icon`
	if len(files) == 0 {
		// log.Println("Recycle bin is empty")
		systray.SetTemplateIcon(icons.EmptyNineTrash, icons.EmptyNineTrash)
	} else if len(files) > 0 {
		systray.SetTemplateIcon(icons.FullNineTrash, icons.FullNineTrash)
		// log.Println("Recycle bin is not empty")
	}
}
