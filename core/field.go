package core

type PointF struct {
	x, y float64
}

type Point struct {
	x, y int
}

type Topology int

const (
	Sqrt3 = 1.73205080756888
)

const (
	CylinderV Topology = iota // N+ <-> N-
	CylinderH                 // E+ <-> W-
	Torus                     // N+ <-> S-, E+ <-> W-
	Sphere                    // N+W+ <-> N-W+
	Rectangle                 // no border pass
)

type Node interface {
	Neighbours() []Node // list of nodes with distance of 1 from this node

	// following nodes are not necessary neighbours, just nodes in the general direction
	North() Node
	South() Node
	East() Node
	West() Node

	Cartesian() (float64, float64) // cartesian corridinates of center
	Logical() (int, int)           // logical coordinates of node
	Data() interface{}
	SetData(interface{})
}

type Field interface {
	AtCartesian(x, y float64) Node
	AtLogical(x, y int) Node
	Topology() Topology
}
