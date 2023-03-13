package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/ikaros/core"
	"github.com/vanilla-os/orchid/cmdr"
)

func NewListDevicesCmd() *cmdr.Command {
	cmd := cmdr.NewCommand(
		"list-devices",
		ikaros.Trans("listDevices.long"),
		ikaros.Trans("listDevices.short"),
		listDevices,
	).WithBoolFlag(
		cmdr.NewBoolFlag(
			"json",
			"j",
			ikaros.Trans("listDevices.json"),
			false,
		),
	)
	cmd.Example = "ikaros list-devices"

	return cmd
}

func listDevices(cmd *cobra.Command, args []string) error {
	if cmd.Flag("json").Changed {
		devicesMap := core.DriversManager{}.GetDevicesAsJson()
		fmt.Println(devicesMap)
		return nil
	}

	spinner, _ := cmdr.Spinner.Start(ikaros.Trans("listDevices.startListing"))
	devicesMap := core.DriversManager{}.GetDevices()
	spinner.Success(ikaros.Trans("listDevices.successfulListing"))

	for group, devices := range devicesMap {
		fmt.Printf(strings.Title(group) + " devices:\n")
		for _, device := range devices {
			fmt.Printf("- ID: %s\n", device.ID)
			fmt.Printf("  Product: %s\n", device.Product)
			fmt.Printf("  Vendor: %s\n", device.Vendor)
			fmt.Printf("  Businfo: %s\n", device.Businfo)
			fmt.Printf("  Drivers: %s\n", strings.Join(device.Drivers, ", "))
		}
		fmt.Println()
	}

	return nil
}
