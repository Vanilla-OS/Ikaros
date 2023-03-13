package core

import (
	"fmt"
)

type PackageManager interface {
	ListDrivers(device Device) []string
	InstallDriver(device Device) error
}

func GetPackageManager() PackageManager {
	pkgManager := GetPkgManager()
	switch pkgManager {
	case "apt":
		return NewAptPackageManager()
	}
	return nil
}

func PkgInstallDriver(device Device) error {
	if GetDebug() {
		fmt.Printf("Fake installing driver for device %s\n", device.ID)
		return nil
	}

	pkgManager := GetPackageManager()
	if pkgManager == nil {
		return fmt.Errorf("no package manager found")
	}
	return pkgManager.InstallDriver(device)
}

func PkgListDrivers(device Device) []string {
	pkgManager := GetPackageManager()
	if pkgManager == nil {
		return nil
	}
	return pkgManager.ListDrivers(device)
}
