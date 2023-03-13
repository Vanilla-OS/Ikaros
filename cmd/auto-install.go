package cmd

import (
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
	cmd.Args = cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)
	cmd.Example = "ikaros auto-install"

	return cmd
}

func autoInstall(cmd *cobra.Command, args []string) error {
	spinner, _ := cmdr.Spinner.Start(ikaros.Trans("autoInstall.startInstallation"))
	err := core.DriversManager{}.AutoInstallDrivers()
	if err != nil {
		spinner.Fail(ikaros.Trans("autoInstall.failedInstallation"))
		return err
	}
	spinner.Success(ikaros.Trans("autoInstall.successfulInstallation"))
	return nil
}
