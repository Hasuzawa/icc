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
