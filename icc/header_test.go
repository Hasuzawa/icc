package icc_test

import (
	"testing"
	"unsafe"

	"github.com/Hasuzawa/icc/icc"
	"github.com/stretchr/testify/assert"
)

func TestHeaderSize(t *testing.T) {
	header := icc.Header{}
	assert.Equal(t, uintptr(128), unsafe.Sizeof(header))
}

func TestHeaderVersion(t *testing.T) {
	for _, tt := range []struct {
		name    string
		value   uint32
		version string
	}{
		{
			name:    "v0.0",
			value:   0x0000_0000,
			version: "0.0.0.0",
		},
		{
			name:    "v1.0",
			value:   0x0100_0000,
			version: "1.0.0.0",
		},
		{
			name:    "v2.3",
			value:   0x0203_0000,
			version: "2.3.0.0",
		},
		{
			name:    "v4.4",
			value:   0x0404_0000,
			version: "4.4.0.0",
		},
		{
			name:    "v12.15.0.0",
			value:   0x0C0F37BA,
			version: "12.15.0.0",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := icc.Header{}
			h.Version = tt.value
			assert.Equal(t, tt.version, h.VersionValue())
		})
	}
}

func TestHeaderDeviceClass(t *testing.T) {
	for _, tt := range []struct {
		name        string
		value       uint32
		deviceClass string
	}{
		{
			name:        "Device Class Input",
			value:       0x7363_6E72,
			deviceClass: "Input",
		},
		{
			name:        "Device Class Display",
			value:       0x6D6E_7472,
			deviceClass: "Display",
		},
		{
			name:        "Device Class Output",
			value:       0x7072_7472,
			deviceClass: "Output",
		},
		{
			name:        "Device Class DeviceLink",
			value:       0x6C69_6E6B,
			deviceClass: "DeviceLink",
		},
		{
			name:        "Device Class ColorSpace",
			value:       0x73706163,
			deviceClass: "ColorSpace",
		},
		{
			name:        "Device Class Abstract",
			value:       0x6162_7374,
			deviceClass: "Abstract",
		},
		{
			name:        "Device Class NamedColor",
			value:       0x6E6D636C,
			deviceClass: "NamedColor",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := icc.Header{}
			h.DeviceClass = tt.value
			assert.Equal(t, tt.deviceClass, h.DeviceClassValue())
		})
	}
}

func TestHeaderColorspace(t *testing.T) {
	for _, tt := range []struct {
		name       string
		value      uint32
		colorSpace string
	}{
		{
			name:       "ColorSpace nCIEXYZ or PCSXYZ",
			value:      0x5859_5A20,
			colorSpace: "nCIEXYZ or PCSXYZ",
		},
		{
			name:       "ColorSpace CIELAB or PCSLAB",
			value:      0x4C61_6220,
			colorSpace: "CIELAB or PCSLAB",
		},
		{
			name:       "ColorSpace CIELUV",
			value:      0x4C75_7620,
			colorSpace: "CIELUV",
		},
		{
			name:       "ColorSpace YCbCr",
			value:      0x5943_6272,
			colorSpace: "YCbCr",
		},
		{
			name:       "ColorSpace CIEYxy",
			value:      0x5978_7920,
			colorSpace: "CIEYxy",
		},
		{
			name:       "ColorSpace RGB",
			value:      0x5247_4220,
			colorSpace: "RGB",
		},
		{
			name:       "ColorSpace Gray",
			value:      0x4752_4159,
			colorSpace: "Gray",
		},
		{
			name:       "ColorSpace HSV",
			value:      0x4853_5620,
			colorSpace: "HSV",
		},
		{
			name:       "ColorSpace HLS",
			value:      0x484C_5320,
			colorSpace: "HLS",
		},
		{
			name:       "ColorSpace CMYK",
			value:      0x434D_594B,
			colorSpace: "CMYK",
		},
		{
			name:       "ColorSpace CMY",
			value:      0x434D_5920,
			colorSpace: "CMY",
		},
		{
			name:       "ColorSpace 2 color",
			value:      0x3243_4C52,
			colorSpace: "2 color",
		},
		{
			name:       "ColorSpace 3 color",
			value:      0x3343_4C52,
			colorSpace: "3 color",
		},
		{
			name:       "ColorSpace 4 color",
			value:      0x3443_4C52,
			colorSpace: "4 color",
		},
		{
			name:       "ColorSpace 5 color",
			value:      0x3543_4C52,
			colorSpace: "5 color",
		},
		{
			name:       "ColorSpace 6 color",
			value:      0x3643_4C52,
			colorSpace: "6 color",
		},
		{
			name:       "ColorSpace 7 color",
			value:      0x3743_4C52,
			colorSpace: "7 color",
		},
		{
			name:       "ColorSpace 8 color",
			value:      0x3843_4C52,
			colorSpace: "8 color",
		},
		{
			name:       "ColorSpace 9 color",
			value:      0x3943_4C52,
			colorSpace: "9 color",
		},
		{
			name:       "ColorSpace 10 color",
			value:      0x4143_4C52,
			colorSpace: "10 color",
		},
		{
			name:       "ColorSpace 11 color",
			value:      0x4243_4C52,
			colorSpace: "11 color",
		},
		{
			name:       "ColorSpace 12 color",
			value:      0x4343_4C52,
			colorSpace: "12 color",
		},
		{
			name:       "ColorSpace 13 color",
			value:      0x4443_4C52,
			colorSpace: "13 color",
		},
		{
			name:       "ColorSpace 14 color",
			value:      0x4543_4C52,
			colorSpace: "14 color",
		},
		{
			name:       "ColorSpace 15 color",
			value:      0x4643_4C52,
			colorSpace: "15 color",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := icc.Header{}
			h.ColorSpace = tt.value
			assert.Equal(t, tt.colorSpace, h.ColorSpaceValue())
		})
	}
}
