package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/vanilla-os/ikaros/core"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewAutoInstallCmd() *cmdr.Command {
	cmd := cmdr.NewCommand(
		"auto-install",
		ikaros.Trans("autoInstall.long"),
		ikaros.Trans("autoInstall.short"),
		autoInstall,
	)
	cmd.Flags().BoolP("list-only", "l", false, ikaros.Trans("autoInstall.listOnly"))
	cmd.Args = cobra.MatchAll(cobra.ExactArgs(0), cobra.OnlyValidArgs)
	cmd.Example = "ikaros auto-install"

	return cmd
}

func autoInstall(cmd *cobra.Command, args []string) error {
	listonly, _ := cmd.Flags().GetBool("list-only")
	var spinner *pterm.SpinnerPrinter
	if !listonly {
		spinner, _ = cmdr.Spinner.Start(ikaros.Trans("autoInstall.startInstallation"))
	}
	err := core.DriversManager{}.AutoInstallDrivers(listonly)
	if spinner == nil {
		return nil
	}
	if err != nil {
		spinner.Fail(ikaros.Trans("autoInstall.failedInstallation"))
		return err
	}
	spinner.Success(ikaros.Trans("autoInstall.successfulInstallation"))
	return nil
}
