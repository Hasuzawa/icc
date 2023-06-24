package header_test

import (
	"encoding/binary"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"

	"github.com/Hasuzawa/icc/header"
	"github.com/Hasuzawa/icc/header/media"
)

func TestHeaderSize(t *testing.T) {
	header := header.Header{}
	assert.Equal(t, uintptr(128), unsafe.Sizeof(header))
}

func TestHeaderCMM(t *testing.T) {
	for _, tt := range []struct {
		name  string
		value string
		cmm   string
	}{
		{
			name:  "Adobe",
			value: "ADBE",
			cmm:   "Adobe Systems, Inc.",
		},
		{
			name:  "Apple",
			value: "appl",
			cmm:   "Apple Computer",
		},
		{
			name:  "Canon",
			value: "CCMS",
			cmm:   "Canon",
		},
		{
			name:  "EFI",
			value: "EFI ",
			cmm:   "EFI",
		},
		{
			name:  "Fujifilm",
			value: "FF  ",
			cmm:   "Fujifilm Corporation",
		},
		{
			name:  "Global Graphics",
			value: "HCMM",
			cmm:   "Global Graphics Software Inc",
		},
		{
			name:  "Canon",
			value: "CCMS",
			cmm:   "Canon",
		},
		{
			name:  "Heidelberger Druckmaschinen AG",
			value: "HDM ",
			cmm:   "Heidelberger Druckmaschinen AG",
		},
		{
			name:  "Hewlett Packard",
			value: "lcms",
			cmm:   "Hewlett Packard",
		},
		{
			name:  "Kodak",
			value: "KCMS",
			cmm:   "Kodak",
		},
		{
			name:  "Konica Minolta",
			value: "MCML",
			cmm:   "Konica Minolta",
		},
		{
			name:  "Microsoft",
			value: "WCS ",
			cmm:   "Microsoft",
		},
		{
			name:  "Toshiba",
			value: "TCMM",
			cmm:   "Toshiba TEC Corporation",
		},
		{
			name:  "Vivo",
			value: "vivo",
			cmm:   "Vivo Mobile Communication",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := header.Header{}
			h.CMM = binary.BigEndian.Uint32([]byte(tt.value))
			assert.Equal(t, tt.cmm, h.CMMValue())
		})
	}
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
			h := header.Header{}
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
			h := header.Header{}
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
			h := header.Header{}
			h.ColorSpace = tt.value
			assert.Equal(t, tt.colorSpace, h.ColorSpaceValue())
		})
	}
}

func TestHeaderDatetime(t *testing.T) {
	for _, tt := range []struct {
		name     string
		value    [12]byte
		datetime time.Time
	}{
		{
			name:     "zero datetime",
			value:    [12]byte{0x0},
			datetime: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "datetime 2013-07-25 18:31",
			value:    [12]byte{0x07, 0xDD, 0x00, 0x07, 0x00, 0x19, 0x00, 0x12, 0x00, 0x1F},
			datetime: time.Date(2013, 7, 25, 18, 31, 0, 0, time.UTC),
		},
		{
			name:     "datetime 1998-10-31 09:41:21",
			value:    [12]byte{0x07, 0xCE, 0x00, 0x0A, 0x00, 0x1F, 0x00, 0x09, 0x00, 0x29, 0x00, 0x15},
			datetime: time.Date(1998, 10, 31, 9, 41, 21, 0, time.UTC),
		},
		{
			name:     "datetime 2010-02-29 (leap year)",
			value:    [12]byte{0x07, 0xDA, 0x00, 0x02, 0x00, 0x1D},
			datetime: time.Date(2010, 2, 29, 0, 0, 0, 0, time.UTC),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := header.Header{}
			h.Datetime = tt.value
			assert.Equal(t, tt.datetime, h.DatetimeValue())
		})
	}
}

func TestHeaderPlatform(t *testing.T) {
	for _, tt := range []struct {
		name     string
		value    uint32
		platform string
	}{
		{
			name:     "Apple Computer, Inc.",
			value:    0x4150_504C,
			platform: "Apple Computer, Inc.",
		},
		{
			name:     "Microsoft Corporation",
			value:    0x4D53_4654,
			platform: "Microsoft Corporation",
		},
		{
			name:     "Silicon Graphics, Inc.",
			value:    0x5347_4920,
			platform: "Silicon Graphics, Inc.",
		},
		{
			name:     "Sun Microsystems, Inc.",
			value:    0x5355_4E57,
			platform: "Sun Microsystems, Inc.",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := header.Header{}
			h.Platform = tt.value
			assert.Equal(t, tt.platform, h.PlatformValue())
		})
	}
}

func TestHeaderFlags(t *testing.T) {
	for _, tt := range []struct {
		name      string
		value     uint32
		embedded  bool
		dependent bool
	}{
		{
			name:      "zero value",
			value:     0b0,
			embedded:  false,
			dependent: false,
		},
		{
			name:      "is embedded",
			value:     0b1000_0000_0000_0000,
			embedded:  true,
			dependent: false,
		},
		{
			name:      "is dependent",
			value:     0b0100_0000_0000_0000,
			embedded:  false,
			dependent: true,
		},
		{
			name:      "is embedded and dependent",
			value:     0b1100_0000_0000_0000,
			embedded:  true,
			dependent: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := header.Header{}
			h.Flags = tt.value
			assert.Equal(t, tt.embedded, h.IsEmbedded())
			assert.Equal(t, tt.dependent, h.IsDependent())
		})
	}
}

func TestHeaderMedia(t *testing.T) {
	for _, tt := range []struct {
		name  string
		value [8]byte
		media media.Media
	}{
		{
			name:  "all zero",
			value: [8]byte{},
			media: media.Media{
				LightMode: false,
				Finish:    false,
				Polarity:  false,
				Vendor:    [4]byte{},
			},
		},
		{
			name:  "flags set to 1",
			value: [8]byte{0b1110_0000},
			media: media.Media{
				LightMode: true,
				Finish:    true,
				Polarity:  true,
				Vendor:    [4]byte{},
			},
		},
		{
			name:  "check vendor bits",
			value: [8]byte{0x0, 0x0, 0x0, 0x0, 0x12, 0x34, 0x56, 0x78},
			media: media.Media{
				LightMode: false,
				Finish:    false,
				Polarity:  false,
				Vendor:    [4]byte{0x12, 0x34, 0x56, 0x78},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := header.Header{}
			h.DeviceAttributes = tt.value
			assert.Equal(t, tt.media, h.Media())
		})
	}
}
