package ros

type Header struct {
	Seq     uint32
	Stamp   Time
	FrameId string
}
