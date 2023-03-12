package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewListDriversCmd() *cmdr.Command {
	cmd := cmdr.NewCommand(
		"list-drivers",
		ikaros.Trans("list-drivers.long"),
		ikaros.Trans("list-drivers.short"),
		listDrivers,
	)
	cmd.Args = cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)
	cmd.Example = "ikaros list-drivers"

	return cmd
}

func listDrivers(cmd *cobra.Command, args []string) error {
	fmt.Println("list-drivers called")
	return nil
}
