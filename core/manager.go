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

type DriversManager struct{}

func (m DriversManager) GetLshw(group string) []Device {
	err := m.isValidGroup(group)
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

	var devices []Device
	for _, device := range data {
		product := device["product"].(string)
		vendor := device["vendor"].(string)
		businfo := device["businfo"].(string)
		device := NewDevice(product, vendor, businfo, nil)
		device.Drivers = PkgListDrivers(*device)
		devices = append(devices, *device)
	}
	return devices
}

func (m DriversManager) isValidGroup(group string) error {
	for _, validGroup := range validGroups {
		if group == validGroup {
			return nil
		}
	}
	return fmt.Errorf("invalid group: %s", group)
}

func (m DriversManager) GetNetworkDevices() []Device {
	return m.GetLshw("network")
}

func (m DriversManager) GetCpuDevices() []Device {
	return m.GetLshw("cpu")
}

func (m DriversManager) GetDisplayDevices() []Device {
	return m.GetLshw("display")
}

func (m DriversManager) GetDevices() map[string][]Device {
	return map[string][]Device{
		"display": m.GetDisplayDevices(),
		"network": m.GetNetworkDevices(),
		"cpu":     m.GetCpuDevices(),
	}
}

func (m DriversManager) GetDevicesAsJson() string {
	devicesMap := m.GetDevices()
	res, err := json.Marshal(devicesMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(res)
}

func (m DriversManager) InstallDriver(device Device) error {
	return PkgInstallDriver(device)
}

func (m DriversManager) GetDeviceByID(id string) (Device, error) {
	devicesMap := m.GetDevices()
	for _, devices := range devicesMap {
		for _, device := range devices {
			if device.ID == id {
				return device, nil
			}
		}
	}
	return Device{}, fmt.Errorf("device not found")
}

func (m DriversManager) AutoInstallDrivers(listonly bool) error {
	devicesMap := m.GetDevices()
	for _, devices := range devicesMap {
		for _, device := range devices {
			if listonly {
				drivers := PkgListDrivers(device)
				if len(drivers) > 0 {
					fmt.Printf("%s ", drivers[0])
				}
				continue
			}
			err := m.InstallDriver(device)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
