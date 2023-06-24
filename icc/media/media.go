package media

type Media struct {
	LightMode bool
	Finish    bool
	Polarity  bool
	Vendor    [4]byte
}

func (m Media) IsReflective() bool {
	return !m.LightMode
}

func (m Media) IsTransprent() bool {
	return m.LightMode
}

func (m Media) IsGlossy() bool {
	return !m.Finish
}

func (m Media) IsMatte() bool {
	return m.Finish
}

func (m Media) IsPositive() bool {
	return !m.Polarity
}

func (m Media) IsNegative() bool {
	return m.Polarity
}
