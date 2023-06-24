package device_class

var (
	deviceClassMap = map[uint32]string{
		0x73636E72: "Input",
		0x6D6E7472: "Display",
		0x70727472: "Output",
		0x6C696E6B: "DeviceLink",
		0x73706163: "ColorSpace",
		0x61627374: "Abstract",
		0x6E6D636C: "NamedColor",
	}
)

func FindDeviceClassBySignature(signature uint32) (deviceClass string, found bool) {
	v, ok := deviceClassMap[signature]
	return v, ok
}
