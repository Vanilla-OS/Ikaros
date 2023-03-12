package cmd

import (
	"embed"

	"github.com/vanilla-os/orchid/cmdr"
)

var ikaros *cmdr.App

func New(version string, fs embed.FS) *cmdr.App {
	ikaros = cmdr.NewApp("ikaros", version, fs)
	return ikaros
}
func NewRootCommand(version string) *cmdr.Command {
	root := cmdr.NewCommand(
		ikaros.Trans("ikaros.use"),
		ikaros.Trans("ikaros.long"),
		ikaros.Trans("ikaros.short"),
		nil)
	root.Version = version

	return root
}
