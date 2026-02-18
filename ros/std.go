package ros

type Header struct {
	Stamp   Time
	FrameId string
}

type SetBoolRequest struct {
	Data bool
}

type SetBoolResponse struct {
	Success bool
	Message string
}
