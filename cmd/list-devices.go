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
		ikaros.Trans("list-devices.long"),
		ikaros.Trans("list-devices.short"),
		listDevices,
	).WithBoolFlag(
		cmdr.NewBoolFlag(
			"json",
			"j",
			ikaros.Trans("list-devices.json"),
			false,
		),
	)
	cmd.Example = "ikaros list-devices"

	return cmd
}

func listDevices(cmd *cobra.Command, args []string) error {
	if cmd.Flag("json").Changed {
		devicesMap := core.LshwParser{}.GetDevicesAsJson()
		fmt.Println(devicesMap)
		return nil
	}

	devicesMap := core.LshwParser{}.GetDevices()
	for group, devices := range devicesMap {
		fmt.Printf(strings.Title(group) + " devices:\n")
		for _, device := range devices {
			fmt.Printf("- %s\n\tVendor: %s\n\tBusinfo: %s\n", device.Product, device.Vendor, device.Businfo)
		}
		fmt.Println()
	}

	return nil
}
