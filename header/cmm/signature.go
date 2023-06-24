package cmm

var (
	CMMMap = map[uint32]string{
		0x41444245: "Adobe Sytems, Inc.",
		0x41434D53: "Agfa Graphics N. V.",
		0x6170706C: "Apple Computer",
		0x43434D53: "Canon",
		0x5543434D: "Canon",
		0x55434D53: "Cannon",
		0x45464920: "EFI ",
		0x46462020: "Fujifilm Corporation",
		0x45584143: "ExactCODE GmbH",
		0x48434d4d: "Global Graphics Software Inc",
		0x6172676C: "Graeme Gill",
		0x44676f53: "GretagMacbeth",
		0x48444d20: "Heidelberger Druckmaschinen AG",
		0x6C636d73: "Hwelett Packard",
		0x52494d58: "ICC",
		0x44494d58: "ICC",
		0x4b434d53: "Kodak",
		0x4d434d53: "Konica Minotla",
		0x57435320: "Microsoft",
		0x5349474E: "Mutoh",
		0x4f4e5958: "Onyx Graphics",
		0x52474d53: "Rolf Gierling Multitools",
		0x53494343: "SampleICC",
		0x54434d4d: "Toshiba TEC Corporation",
		0x33324254: "the imaging factory",
		0x7669766f: "Vivo Mobile Communication",
		0x57544720: "Ware To Go",
		0x7a633030: "Zoran Corporation",
	}
)

func FindCMMBySignature(signature uint32) (cmm string, found bool) {
	v, ok := CMMMap[signature]
	return v, ok
}
