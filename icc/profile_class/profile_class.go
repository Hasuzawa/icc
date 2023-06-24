package profile_class

var (
	profileClassMap = map[uint32]string{
		0x73636E72: "Input",
		0x6D6E7472: "Display",
		0x70727472: "Output",
		0x6C696E6B: "DeviceLink",
		0x73706163: "ColorSpace",
		0x61627374: "Abstract",
		0x6E6D636C: "NamedColor",
	}
)

func FindProfileClassById(id uint32) (profileClass string, found bool) {
	v, ok := profileClassMap[id]
	return v, ok
}
