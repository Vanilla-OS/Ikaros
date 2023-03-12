package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewAutoInstallCmd() *cmdr.Command {
	cmd := cmdr.NewCommand(
		"auto-install",
		ikaros.Trans("auto-install.long"),
		ikaros.Trans("auto-install.short"),
		autoInstall,
	)
	cmd.Args = cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)
	cmd.Example = "ikaros auto-install"

	return cmd
}

func autoInstall(cmd *cobra.Command, args []string) error {
	fmt.Println("auto-install called")
	return nil
}
