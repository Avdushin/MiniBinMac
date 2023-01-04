package main

import (
	"MiniBinMac/src/pkg/ui"
	"MiniBinMac/src/pkg/ui/layout"
	"MiniBinMac/src/utils"

	"github.com/getlantern/systray"
)

func main() {
	// logging
	utils.LogsInit("logs.log")
	// Get home directory
	go utils.GetHomeDir()
	// Run system tray application
	systray.Run(ui.RenderSysTray, layout.OnExit)
}
