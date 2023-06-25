package header

import (
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"github.com/Hasuzawa/icc/header/cmm"
	"github.com/Hasuzawa/icc/header/color_space"
	"github.com/Hasuzawa/icc/header/device_class"
	"github.com/Hasuzawa/icc/header/manufacturer"
	"github.com/Hasuzawa/icc/header/media"
	"github.com/Hasuzawa/icc/header/platform"
	"github.com/Hasuzawa/icc/header/rendering_intent"
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
	if h.DeviceClassValue() != "DeviceLink" {
		fmt.Println(h.PCSValue())
		if !strings.Contains(h.PCSValue(), "PCSXYZ") && !strings.Contains(h.PCSValue(), "PCSLAB") {
			return ErrInvalidPCSField
		}
	}
	return nil
}

func (h Header) CMMValue() string {
	cmm, found := cmm.FindCMMBySignature(h.CMM)
	if found {
		return cmm
	}
	return ""
}

func (h Header) VersionValue() string {
	var (
		major uint8 = uint8(h.Version >> 24)
		minor uint8 = uint8((h.Version >> 16) & 0xff)
	)
	return fmt.Sprintf("%v.%v.%v.%v", major, minor, 0, 0)
}

func (h Header) DeviceClassValue() string {
	deviceClass, found := device_class.FindDeviceClassBySignature(h.DeviceClass)
	if found {
		return deviceClass
	}
	return ""
}

func (h Header) ColorSpaceValue() string {
	colorSpace, found := color_space.FindColorSpaceBySignature(h.ColorSpace)
	if found {
		return colorSpace
	}
	return ""
}

func (h Header) PCSValue() string {
	colorSpace, found := color_space.FindColorSpaceBySignature(h.PCS)
	if found {
		return colorSpace
	}
	return ""
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

func (h Header) PlatformValue() string {
	platform, found := platform.FindPlatformBySignature(h.Platform)
	if found {
		return platform
	}
	return ""
}

func (h Header) IsEmbedded() bool {
	return (h.Flags & 0b1000_0000_0000_0000) != 0b0
}

func (h Header) IsDependent() bool {
	return (h.Flags & 0b0100_0000_0000_0000) != 0b0
}

func (h Header) Media() media.Media {
	a := h.DeviceAttributes
	var vendorBits [4]byte
	copy(vendorBits[:], a[4:])
	return media.Media{
		LightMode: (a[0] & 0b1000_0000) != 0b0,
		Finish:    (a[0] & 0b0100_0000) != 0b0,
		Polarity:  (a[0] & 0b0010_0000) != 0b0,
		Vendor:    vendorBits,
	}
}

func (h Header) RenderingIntentValue() string {
	renderingIntent, found := rendering_intent.FindRenderingIntentByValue(h.RenderingIntent)
	if found {
		return renderingIntent
	}
	return ""
}

func (h Header) ManufacturerName() string {
	name, found := manufacturer.FindManufacturerBySignature(h.DeviceManufacturer)
	if found {
		return name
	}
	return ""
}
