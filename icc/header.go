package icc

// The profile header provides the necessary information to allow a receiving system to properly search and sort
// ICC profiles. The profile header is 128 bytes in length and contains 18 fields.
type Header struct {
	Size               uint32
	CMM                uint32
	Version            uint32
	DeviceClass        uint32
	ColorSpace         uint32
	PCS                [4]byte
	Datetime           [12]byte
	Signature          [4]byte
	Flags              [4]byte
	DeviceManufacturer uint32
	DeviceModel        uint32
	DeviceAttributes   [8]byte
	RenderingIntent    uint32
	PCSIlluminant      [12]byte
	Creator            uint32
	ID                 [16]byte
	Reserved           [28]byte
}
