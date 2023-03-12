package core

import (
	"fmt"
)

type PackageManager interface {
	ListDrivers(device Device) []string
	InstallDriver(driver string) error
}

func GetPackageManager() PackageManager {
	pkgManager := GetPkgManager()
	switch pkgManager {
	case "apt":
		return NewAptPackageManager()
	}
	return nil
}

func InstallDriver(driver string) error {
	pkgManager := GetPackageManager()
	if pkgManager == nil {
		return fmt.Errorf("no package manager found")
	}
	return pkgManager.InstallDriver(driver)
}

func ListDrivers(device Device) []string {
	pkgManager := GetPackageManager()
	if pkgManager == nil {
		return nil
	}
	return pkgManager.ListDrivers(device)
}
