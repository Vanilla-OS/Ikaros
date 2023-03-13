package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vanilla-os/ikaros/core"
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
	spinner, _ := cmdr.Spinner.Start(ikaros.Trans("auto-install.startInstallation"))
	err := core.DriversManager{}.AutoInstallDrivers()
	if err != nil {
		spinner.Fail(ikaros.Trans("auto-install.failedInstallation"))
		return err
	}
	spinner.Success(ikaros.Trans("auto-install.successfulInstallation"))
	return nil
}
