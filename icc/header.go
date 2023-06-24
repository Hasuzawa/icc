package icc

import (
	"fmt"

	"github.com/Hasuzawa/icc/icc/device_class"
	"github.com/Hasuzawa/icc/icc/manufacturer"
)

// The profile header provides the necessary information to allow a receiving system to properly search and sort
// ICC profiles. The profile header is 128 bytes in length and contains 18 fields.
type Header struct {
	Size               uint32
	CMM                uint32
	Version            uint32
	DeviceClass        uint32
	ColorSpace         uint32
	PCS                uint32
	Datetime           [12]byte
	Signature          uint32
	Platform           uint32
	Flags              uint32
	DeviceManufacturer uint32
	DeviceModel        uint32
	DeviceAttributes   [8]byte
	RenderingIntent    uint32
	PCSIlluminant      [12]byte
	Creator            uint32
	Id                 [16]byte
	Reserved           [28]byte
}

func (h Header) Validate() error {
	return nil
}

func (h Header) VersionValue() string {
	var (
		major uint8 = uint8(h.Version >> 24)
		minor uint8 = uint8((h.Version >> 16) & 0xff)
	)
	return fmt.Sprintf("%v.%v.%v.%v", major, minor, 0, 0)
}

func (h Header) DeviceClassValue() string {
	deviceClass, found := device_class.FindDeviceClassById(h.DeviceClass)
	if found {
		return deviceClass
	}
	return "unknown"
}

func (h Header) ManufacturerName() string {
	name, found := manufacturer.FindManufacturerById(h.DeviceManufacturer)
	if found {
		return name
	}
	return "unknown"
}
