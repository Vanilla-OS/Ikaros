package core

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

type AptPackageManager struct {
	driverPrefixes map[string]string
}

func NewAptPackageManager() *AptPackageManager {
	return &AptPackageManager{
		driverPrefixes: map[string]string{
			// this list can be used to simplify the search, if we know
			// the driver name is always the same for a given device
			// writing it here avoids running apt-cache search
			// based on the device name
			"GeForce GTX": "nvidia-driver",
			"GeForce RTX": "nvidia-driver",
		},
	}
}

func (a AptPackageManager) ListDrivers(device Device) []string {
	drivers := []string{}

	for prefix, pattern := range a.driverPrefixes {
		if strings.Contains(device.Product, prefix) {
			drivers = a.listPackagesByBusinfo(pattern, device.Businfo)
			break
		}
	}

	if len(drivers) == 0 {
		// the following needs to be improved, it's just a quick hack
		for _, word := range strings.Fields(device.Product)[:2] {
			if !regexp.MustCompile(`^[a-zA-Z0-9-]{2,}$`).MatchString(word) {
				continue
			}

			pattern := strings.ToLower(word)
			if packages := a.listPackagesByBusinfo(pattern, device.Businfo); packages != nil {
				drivers = append(drivers, packages...)
			}
		}
	}

	sort.Slice(drivers, func(i, j int) bool {
		return drivers[i] > drivers[j]
	})

	return drivers
}

func (a *AptPackageManager) InstallDriver(driver string) error {
	cmd := exec.Command("sudo", "apt-get", "install", driver)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (a AptPackageManager) listPackagesByBusinfo(pattern, businfo string) []string {
	cmd := exec.Command("apt-cache", "search", pattern)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(&out)
	packages := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, pattern) {
			packages = append(packages, strings.Split(line, " ")[0])
		}
	}

	drivers := []string{}
	for _, pkg := range packages {
		modaliases := a.getPackageModaliases(pkg)
		if a.matchBusinfo(modaliases, businfo) {
			drivers = append(drivers, pkg)
		}
	}

	return drivers
}

func (a AptPackageManager) listPackages(pattern string) []string {
	cmd := exec.Command("apt-cache", "search", pattern)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(&out)
	packages := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, pattern) {
			packages = append(packages, strings.Split(line, " ")[0])
		}
	}

	return packages
}

func (a AptPackageManager) getPackageModaliases(pkg string) string {
	cmd := exec.Command("apt-cache", "show", pkg)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	res := string(out)
	scanner := bufio.NewScanner(strings.NewReader(res))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Modaliases:") {
			modaliases := strings.TrimSpace(strings.TrimPrefix(line, "Modaliases:"))
			return modaliases
		}
	}
	return ""
}

func (a AptPackageManager) matchBusinfo(modalias, businfo string) bool {
	if modalias == "" {
		return false
	}

	vendorProductRe := regexp.MustCompile(`pci:v0000(.+)d0000(.+)sv`)
	vendorProduct := vendorProductRe.FindStringSubmatch(modalias)
	if vendorProduct == nil {
		return false
	}

	vendor := vendorProduct[1]
	product := vendorProduct[2]

	businfoRe := regexp.MustCompile(`pci@0000:(.+?):(.+).0`)
	businfo = businfoRe.FindStringSubmatch(businfo)[0]
	if businfo == "" {
		return false
	}

	bus := string(businfo[1])
	slot := string(businfo[2])

	return strings.Contains(
		strings.ToLower(modalias),
		strings.ToLower(vendor),
	) && strings.Contains(
		strings.ToLower(modalias),
		strings.ToLower(product),
	) && strings.Contains(
		strings.ToLower(modalias),
		strings.ToLower(bus),
	) && strings.Contains(
		strings.ToLower(modalias),
		strings.ToLower(slot),
	)
}
