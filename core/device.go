package core

import (
	"crypto/sha256"
	"fmt"
)

type Device struct {
	Product string
	Vendor  string
	Businfo string
	Drivers []string
	ID      string
}

func NewDevice(product string, vendor string, businfo string, drivers []string) *Device {
	id := generateDeviceID(product, vendor, businfo)
	return &Device{
		Product: product,
		Vendor:  vendor,
		Businfo: businfo,
		Drivers: drivers,
		ID:      id,
	}
}

func generateDeviceID(product string, vendor string, businfo string) string {
	id := fmt.Sprintf("%s_%s_%s", product, vendor, businfo)
	hash := sha256.Sum256([]byte(id))
	return fmt.Sprintf("%x", hash)[:8]
}
