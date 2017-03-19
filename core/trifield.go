package core

import (
	// "fmt"
	"math"
)

/*
 *	struct triNode
 */
type triNode struct {
	field   *triField
	logical Point
	data    interface{}
}

func (n triNode) Data() interface{} {
	return n.data
}

func (n *triNode) SetData(d interface{}) {
	n.data = d
}

func (n triNode) Logical() (int, int) {
	return n.logical.X, n.logical.Y
}

func (n triNode) LogicalP() Point {
	return n.logical
}

func (n triNode) Cartesian() (float64, float64) {
	x := Sqrt3 / 2. * float64(n.logical.X)
	y := 1.5 * float64(n.logical.Y)
	even := (n.logical.X+n.logical.Y)%2 == 0
	if !even {
		y += 0.5
	}
	return x, y
}

func (n triNode) North() Node {
	return n.field.AtLogical(n.logical.X, n.logical.Y-1)
}
func (n triNode) South() Node {
	return n.field.AtLogical(n.logical.X, n.logical.Y+1)
}
func (n triNode) East() Node {
	return n.field.AtLogical(n.logical.X+1, n.logical.Y)
}
func (n triNode) West() Node {
	return n.field.AtLogical(n.logical.X-1, n.logical.Y)
}

func (n triNode) Neighbours() (r []Node) {
	if node := n.East(); node != nil {
		r = append(r, node)
	}
	if node := n.West(); node != nil {
		r = append(r, node)
	}

	var node Node
	even := (n.logical.X+n.logical.Y)%2 == 0
	if even {
		node = n.North()
	} else {
		node = n.South()
	}
	if node != nil {
		r = append(r, node)
	}
	return
}

func (n triNode) Field() Field {
	return n.field
}

/*
 *	struct triField
 */
type triField struct {
	nodes [][]interface{}
}

// check if field contains cell with logical coordinates x, y
func (f *triField) inBounds(x, y int) bool {
	return 0 <= x && x < len(f.nodes) && 0 <= y && y < len(f.nodes[0])
}

func (f *triField) AtCartesian(x, y float64) Node {

	xstr := math.Floor(x * 2. / Sqrt3) // stripe number
	ystr := math.Floor(2. / 3. * (y + 0.5))

	// fmt.Printf("xstr: %v ystr: %v\n", xstr, ystr)

	/*
		      0 1 2 3 4 5 6
			----------------
		  0 \  /\  /\  /\  /
			 \/  \/  \/  \/
			----------------
		  1  /\  /\  /\  /\
		    /  \/  \/  \/  \
		    ----------------
	*/
	var x1, y1, x2, y2 float64 // line that divides current rectangle
	//bottom left to upper right
	x1 = xstr * Sqrt3 / 2.
	y1 = (ystr+1.)*1.5 - 0.5
	x2 = (xstr + 1.) * Sqrt3 / 2.
	y2 = ystr*1.5 - 0.5
	even := (int(xstr)+int(ystr))%2 == 0
	if !even {
		//upper left to bottom right
		y1, y2 = y2, y1
	}

	upper := (y1-y2)*(x1-x) > (x1-x2)*(y1-y)

	logX := int(xstr)
	if even != upper {
		logX += 1
	}

	logY := int(ystr)

	if !f.inBounds(logX, logY) {
		return nil
	}

	ret := triNode{
		field:   f,
		logical: Point{logX, logY},
		data:    f.nodes[logX][logY]}
	return &ret
}

func (f *triField) AtLogical(x, y int) Node {
	if !f.inBounds(x, y) {
		return nil
	}

	ret := triNode{
		field:   f,
		logical: Point{x, y},
		data:    f.nodes[x][y]}
	return &ret
}

func (triField) Topology() Topology {
	return Rectangle
}

func NewTriField(w, h int) Field {
	ret := triField{}
	ret.nodes = make([][]interface{}, w)
	for i := 0; i < w; i++ {
		ret.nodes[i] = make([]interface{}, h)
	}
	return &ret
}
