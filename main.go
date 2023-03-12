package main

import (
	"embed"

	"github.com/vanilla-os/ikaros/cmd"
	"github.com/vanilla-os/orchid/cmdr"
)

var (
	Version = "0.0.1"
)

//go:embed locales/*.yml
var fs embed.FS
var ikaros *cmdr.App

func main() {
	ikaros = cmd.New(Version, fs)

	// root command
	root := cmd.NewRootCommand(Version)
	ikaros.CreateRootCommand(root)

	listDevicesCmd := cmd.NewListDevicesCmd()
	root.AddCommand(listDevicesCmd)

	installCmd := cmd.NewInstallCmd()
	root.AddCommand(installCmd)

	autoInstallCmd := cmd.NewAutoInstallCmd()
	root.AddCommand(autoInstallCmd)

	// run the app
	err := ikaros.Run()
	if err != nil {
		cmdr.Error.Println(err)
	}
}
