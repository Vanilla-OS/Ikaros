package cmd

import (
	"fmt"
	"os"

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
	cmd.Flags().BoolP("list-only", "l", false, ikaros.Trans("install.listOnly"))
	cmd.Args = cobra.MinimumNArgs(1)
	cmd.Example = "ikaros install"

	return cmd
}

func install(cmd *cobra.Command, args []string) error {
	listonly, _ := cmd.Flags().GetBool("list-only")
	if listonly {
		device, err := core.DriversManager{}.GetDeviceByID(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, ikaros.Trans("install.failedGetDevice"))
			return err
		}

		drivers := core.PkgListDrivers(device)
		fmt.Printf("%s", drivers[0])
		return nil
	}

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
