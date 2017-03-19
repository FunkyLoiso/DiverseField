package algorithms

import (
	"fmt"
	"github.com/FunkyLoiso/DiverseField/core"
	"reflect"
	"testing"
)

func distance(from, to core.Node) float64 {
	x, _ := to.Cartesian()
	if x > 1.5 {
		return 10.
	}
	return 1.
}

type Field interface {
}

type Node interface {
}

type Point struct {
	x, y int
}

type triField struct {
	nodes [][]interface{}
}

type triNode struct {
	field   *triField
	logical Point
	data    interface{}
}

// check if field contains cell with logical coordinates x, y
func (f *triField) inBounds(x, y int) bool {
	return 0 <= x && x < len(f.nodes) && 0 <= y && y < len(f.nodes[0])
}

func NewTriField(w, h int) Field {
	ret := triField{}
	ret.nodes = make([][]interface{}, w)
	for i := 0; i < w; i++ {
		ret.nodes[i] = make([]interface{}, h)
	}
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

func TestAlgo(t *testing.T) {
	field := core.NewTriField(3, 3)
	from, to := field.AtLogical(0, 0), field.AtLogical(2, 2)
	from2 := field.AtLogical(0, 0)
	fmt.Printf("from  type: %v, value: %v\n", reflect.TypeOf(from), reflect.ValueOf(from))
	fmt.Printf("from2 type: %v, value: %v\n", reflect.TypeOf(from2), reflect.ValueOf(from2))
	fmt.Printf("from2 type: %v, value: %v\n", reflect.TypeOf(from2), reflect.ValueOf(from2))
	fmt.Printf("path from '%v' to '%v'\n", from, to)
	path := AStarShortestPath(from, to, distance)
	fmt.Print("the path: ")
	for _, v := range path {
		fmt.Printf("%v ", v.LogicalP())
	}

	// f := NewTriField(10, 10).(*triField)
	// a, b := f.AtLogical(0, 0).(*triNode), f.AtLogical(0, 0).(*triNode)
	// fmt.Println(a.data == b.data, a.field == b.field, a.logical == b.logical, a == b)

	// c, d := triNode{f, Point{0, 0}, nil}, triNode{f, Point{0, 0}, nil}
	// cp, dp := &c, &d
	// fmt.Println(cp.data == dp.data, cp.field == dp.field, cp.logical == dp.logical, cp == dp)
}
