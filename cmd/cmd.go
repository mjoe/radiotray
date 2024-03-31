// Package cmd implement main cmd interface
package cmd

import (
	"log"

	"github.com/mjoe/radiotray/cmd/config"
	"github.com/mjoe/radiotray/cmd/player"
	"github.com/mjoe/radiotray/cmd/settings"
	"github.com/mjoe/radiotray/cmd/tray"
	"github.com/sqweek/dialog"
)

// Main func call
func Main() {
	// Load config
	cfg := config.New()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	// Init settings
	// settings.Init()

	tray.Init(cfg, func() {
		for {
			select {
			case r := <-tray.RadioSelectCH:
				if err := player.Play(r); err != nil {
					log.Println("ERROR", err)
					continue
				}

				tray.SetPlayingIcon()
				tray.EnableStopItem()
			case <-tray.StopCH:
				player.Stop()

				tray.UncheckRadioItems()
				tray.SetDefaultIcon()
				tray.DisableStopItem()
			case <-tray.SettingsCH:
				settings.Open()
			case <-tray.InfoCH:
				dialog.Message("RadioTray Version %s", config.Version).Title("RadioTray").Info()
			}
		}
	})
}
