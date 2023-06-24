package platform

var (
	platformSignatureMap = map[uint32]string{
		0x4150_504C: "Apple Computer, Inc.",
		0x4D53_4654: "Microsoft Corporation",
		0x5347_4920: "Silicon Graphics, Inc.",
		0x5355_4E57: "Sun Microsystems, Inc.",
	}
)

func FindPlatformBySignature(signature uint32) (platform string, found bool) {
	v, ok := platformSignatureMap[signature]
	return v, ok
}
