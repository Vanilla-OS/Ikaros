package core

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

var validGroups = []string{
	"network",
	"cpu",
	"display",
}

type LshwParser struct{}

func (l LshwParser) GetLshw(group string) []Device {
	err := l.isValidGroup(group)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	cmd := exec.Command("lshw", "-C", group, "-json")
	res, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var data []map[string]interface{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	pkgManager := GetPackageManager()

	var devices []Device
	for _, device := range data {
		product := device["product"].(string)
		vendor := device["vendor"].(string)
		businfo := device["businfo"].(string)
		device := Device{product, vendor, businfo, nil}
		device.Drivers = pkgManager.ListDrivers(device)
		devices = append(devices, device)
	}
	return devices
}

func (l LshwParser) isValidGroup(group string) error {
	for _, validGroup := range validGroups {
		if group == validGroup {
			return nil
		}
	}
	return fmt.Errorf("invalid group: %s", group)
}

func (l LshwParser) GetNetworkDevices() []Device {
	return l.GetLshw("network")
}

func (l LshwParser) GetCpuDevices() []Device {
	return l.GetLshw("cpu")
}

func (l LshwParser) GetDisplayDevices() []Device {
	return l.GetLshw("display")
}

func (l LshwParser) GetDevices() map[string][]Device {
	return map[string][]Device{
		"display": l.GetDisplayDevices(),
		"network": l.GetNetworkDevices(),
		"cpu":     l.GetCpuDevices(),
	}
}

func (l LshwParser) GetDevicesAsJson() string {
	devicesMap := l.GetDevices()
	res, err := json.Marshal(devicesMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(res)
}
