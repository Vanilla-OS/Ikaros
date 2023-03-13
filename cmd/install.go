package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vanilla-os/ikaros/core"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewInstallCmd() *cmdr.Command {
	cmd := cmdr.NewCommand(
		"install",
		ikaros.Trans("install.long"),
		ikaros.Trans("install.short"),
		install,
	)
	cmd.Args = cobra.MinimumNArgs(1)
	cmd.Example = "ikaros install"

	return cmd
}

func install(cmd *cobra.Command, args []string) error {
	spinner, _ := cmdr.Spinner.Start(ikaros.Trans("install.startInstallation"))
	device, err := core.DriversManager{}.GetDeviceByID(args[0])
	if err != nil {
		spinner.Fail(ikaros.Trans("install.failedGetDevice"))
		return err
	}

	err = core.DriversManager{}.InstallDriver(device)
	if err != nil {
		spinner.Fail(ikaros.Trans("install.failedInstallDriver"))
		return err
	}

	spinner.Success(ikaros.Trans("install.successfulInstallation"))
	return nil
}
