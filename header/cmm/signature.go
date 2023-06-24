package cmm

var (
	CMMMap = map[uint32]string{
		0x4144_4245: "Adobe Systems, Inc.",
		0x4143_4D53: "Agfa Graphics N. V.",
		0x6170_706C: "Apple Computer",
		0x4343_4D53: "Canon",
		0x5543_434D: "Canon",
		0x5543_4D53: "Canon",
		0x4546_4920: "EFI",
		0x4646_2020: "Fujifilm Corporation",
		0x4558_4143: "ExactCODE GmbH",
		0x4843_4d4d: "Global Graphics Software Inc",
		0x6172_676C: "Graeme Gill",
		0x4467_6f53: "GretagMacbeth",
		0x4844_4d20: "Heidelberger Druckmaschinen AG",
		0x6C63_6d73: "Hewlett Packard",
		0x5249_4d58: "ICC",
		0x4449_4d58: "ICC",
		0x4b43_4d53: "Kodak",
		0x4d43_4d44: "Konica Minolta", // This value is listed on Tag Registry 4 March 2021 but it is wrong.
		0x4d43_4d4c: "Konica Minolta",
		0x5743_5320: "Microsoft",
		0x5349_474E: "Mutoh",
		0x4f4e_5958: "Onyx Graphics",
		0x5247_4d53: "Rolf Gierling Multitools",
		0x5349_4343: "SampleICC",
		0x5443_4d4d: "Toshiba TEC Corporation",
		0x3332_4254: "the imaging factory",
		0x7669_766f: "Vivo Mobile Communication",
		0x5754_4720: "Ware To Go",
		0x7a63_3030: "Zoran Corporation",
	}
)

func FindCMMBySignature(signature uint32) (cmm string, found bool) {
	v, ok := CMMMap[signature]
	return v, ok
}
