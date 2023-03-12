package core

import "fmt"

type Device struct {
	Product string
	Vendor  string
	Businfo string
}

func (d Device) String() string {
	return fmt.Sprintf("Product: %s, Vendor: %s, Businfo: %s", d.Product, d.Vendor, d.Businfo)
}

func (d Device) AsJson() string {
	return fmt.Sprintf(`{"product": "%s", "vendor": "%s", "businfo": "%s"}`, d.Product, d.Vendor, d.Businfo)
}
