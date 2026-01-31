package ros

type JointState struct {
	Header   Header
	Name     []string
	Position []float64
	Velocity []float64
	Effort   []float64
}

type PointFieldDataType uint8

const (
	PointFieldInt8 PointFieldDataType = iota + 1
	PointFieldUint8
	PointFieldInt16
	PointFieldUint16
	PointFieldInt32
	PointFieldUint32
	PointFieldFloat32
	PointFieldFloat64
)

type PointField struct {
	Name     string
	Offset   uint32
	Datatype PointFieldDataType
	Count    uint32
}

type PointCloud2 struct {
	Header      Header
	Height      uint32
	Width       uint32
	Fields      []PointField
	IsBigEndian bool
	PointStep   uint32
	RowStep     uint32
	Data        []uint8
	IsDense     bool
}
