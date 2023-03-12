package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewInstallCmd() *cmdr.Command {
	cmd := cmdr.NewCommand(
		"install",
		ikaros.Trans("install.long"),
		ikaros.Trans("install.short"),
		install,
	)
	cmd.Args = cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)
	cmd.Example = "ikaros install"

	return cmd
}

func install(cmd *cobra.Command, args []string) error {
	fmt.Println("install called")
	return nil
}
