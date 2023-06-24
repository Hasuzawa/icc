package icc

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/Hasuzawa/icc/icc/color_space"
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

func (h Header) ColorSpaceValue() string {
	colorSpace, found := color_space.FindColorSpaceBySignature(h.ColorSpace)
	if found {
		return colorSpace
	}
	return "unknown"
}

func (h Header) DatetimeValue() time.Time {
	var (
		year   = int(binary.BigEndian.Uint16(h.Datetime[0:2]))
		month  = int(binary.BigEndian.Uint16(h.Datetime[2:4]))
		day    = int(binary.BigEndian.Uint16(h.Datetime[4:6]))
		hour   = int(binary.BigEndian.Uint16(h.Datetime[6:8]))
		minute = int(binary.BigEndian.Uint16(h.Datetime[8:10]))
		second = int(binary.BigEndian.Uint16(h.Datetime[10:12]))
	)
	fmt.Println(int(h.Datetime[0]) << 8)
	fmt.Println(int(h.Datetime[1]))
	return time.Date(
		year,
		time.Month(month),
		day,
		hour,
		minute,
		second,
		0,
		time.UTC,
	)
}

func (h Header) ManufacturerName() string {
	name, found := manufacturer.FindManufacturerById(h.DeviceManufacturer)
	if found {
		return name
	}
	return "unknown"
}
