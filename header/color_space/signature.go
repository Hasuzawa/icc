package color_space

var (
	colorSpaceMap = map[uint32]string{
		0x5859_5A20: "nCIEXYZ or PCSXYZ",
		0x4C61_6220: "CIELAB or PCSLAB",
		0x4C75_7620: "CIELUV",
		0x5943_6272: "YCbCr",
		0x5978_7920: "CIEYxy",
		0x5247_4220: "RGB",
		0x4752_4159: "Gray",
		0x4853_5620: "HSV",
		0x484C_5320: "HLS",
		0x434D_594B: "CMYK",
		0x434D_5920: "CMY",
		0x3243_4C52: "2 color",
		0x3343_4C52: "3 color",
		0x3443_4C52: "4 color",
		0x3543_4C52: "5 color",
		0x3643_4C52: "6 color",
		0x3743_4C52: "7 color",
		0x3843_4C52: "8 color",
		0x3943_4C52: "9 color",
		0x4143_4C52: "10 color",
		0x4243_4C52: "11 color",
		0x4343_4C52: "12 color",
		0x4443_4C52: "13 color",
		0x4543_4C52: "14 color",
		0x4643_4C52: "15 color",
	}
)

func FindColorSpaceBySignature(signature uint32) (signatureName string, found bool) {
	v, ok := colorSpaceMap[signature]
	return v, ok
}
