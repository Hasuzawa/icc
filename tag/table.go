package tag

type TagTable struct {
	Count      uint32
	TagEntries []*TagEntry
}

type TagEntry struct {
	Signature uint32
	Offset    uint32
	Size      uint32
}
