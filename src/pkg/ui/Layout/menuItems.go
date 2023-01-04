package layout

import (
	"log"

	"github.com/getlantern/systray"
)

func SetupMenu() {
	go OnExit()
}

func OnExit() {
	// Exit listener
	mQuit := systray.AddMenuItem("Выход", "Quit the whole app")

	go func() {
		<-mQuit.ClickedCh
		log.Println("Clicked Quit button")
		log.Println("The App was successfully closed.")
		systray.Quit()
	}()
}
