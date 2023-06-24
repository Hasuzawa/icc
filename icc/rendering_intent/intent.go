package rendering_intent

var (
	renderingIntentMap = map[uint32]string{
		0x0000_0000: "Perceptual",
		0x0000_0001: "Media-relative colorimetric",
		0x0000_0002: "Saturation",
		0x0000_0003: "ICC-absolute colorimetric",
	}
)

func FindRenderingIntentByValue(value uint32) (renderingIntent string, found bool) {
	v, ok := renderingIntentMap[value]
	return v, ok
}
