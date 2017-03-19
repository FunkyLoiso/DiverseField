package algorithms

import (
	"container/heap"
	"fmt"
	"github.com/FunkyLoiso/DiverseField/core"
	"math"
	// "time"
)

type DistanceFunc func(from, to core.Node) float64 // distance between from and to assuming that they are neighbours

func costEstimate(from, to core.Node) float64 {
	x1, y1 := from.Cartesian()
	x2, y2 := to.Cartesian()
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

/*
 *	Node priority queue based on estimated cost through this node to finish
 */
type pqNode struct {
	node core.Node
	cost float64
}

type pq struct {
	arr  []*pqNode
	dict map[core.Point]int
}

func (pq pq) Len() int {
	return len(pq.arr)
}

func (pq pq) Less(i, j int) bool {
	return pq.arr[i].cost < pq.arr[j].cost
}

func (pq pq) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
	pq.dict[pq.arr[i].node.LogicalP()] = i
	pq.dict[pq.arr[j].node.LogicalP()] = j
}

func (pq *pq) Push(x interface{}) {
	item := x.(*pqNode)
	pq.dict[item.node.LogicalP()] = len(pq.arr)
	pq.arr = append(pq.arr, item)
}

func (pq *pq) Pop() interface{} {
	item := pq.arr[len(pq.arr)-1]
	pq.arr = pq.arr[:len(pq.arr)-1]
	delete(pq.dict, item.node.LogicalP())
	return item
}

// and an update cost function
func (pq *pq) update(loc core.Point, cost float64) {
	idx, ok := pq.dict[loc]
	if !ok {
		panic(fmt.Sprintf("Location '%v' not found", loc))
	}
	pq.arr[idx].cost = cost
	heap.Fix(pq, idx)
}

func (pq pq) find(loc core.Point) (*pqNode, bool) {
	if idx, ok := pq.dict[loc]; ok {
		return pq.arr[idx], true
	} else {
		return nil, false
	}
}

/*
 *	A* implemetention
 */
func AStarShortestPath(from, to core.Node, distance DistanceFunc) []core.Node {
	// for now only rectangular fields are supported
	if from.Field() != to.Field() {
		panic(fmt.Sprintf("Nodes are from different fileds: '%v' '%x'", from, to))
	}
	if from.Field().Topology() != core.Rectangle {
		panic(fmt.Sprint("Only rectangle fields are supported"))
	}

	evaluated := make(map[core.Point]struct{})
	discovered := pq{arr: make([]*pqNode, 0), dict: make(map[core.Point]int)}
	cameFrom := make(map[core.Point]core.Node)  // from which node can be most effective reached from
	startToCost := make(map[core.Point]float64) // const from start tot node
	// startThroughCost	map[core.Node]float64   // aprox cost from start tot finish through node

	heap.Push(&discovered, &pqNode{node: from, cost: costEstimate(from, to)})
	startToCost[from.LogicalP()] = 0.
	// startThroughCost[from] = costEstimate(from, to)

	for discovered.Len() != 0 {
		// time.Sleep(time.Second)
		current := heap.Pop(&discovered).(*pqNode)
		fmt.Printf("current: '%v', cost: %v, heap: %v\n", current.node, current.cost, discovered.arr)
		if current.node.LogicalP() == to.LogicalP() {
			return reconstructPath(cameFrom, current.node)
		}

		evaluated[current.node.LogicalP()] = struct{}{}

		for _, neighbor := range current.node.Neighbours() {
			fmt.Printf("neibhor: '%v'\n", neighbor)
			if _, found := evaluated[neighbor.LogicalP()]; found {
				continue
			}

			tentStartToCost := current.cost + distance(current.node, neighbor)
			if pqNeighbor, found := discovered.find(neighbor.LogicalP()); !found {
				// not yet discovered, add with current cost + distance
				pqNeighbor = &pqNode{node: neighbor, cost: tentStartToCost + costEstimate(neighbor, to)}
				heap.Push(&discovered, pqNeighbor)
			} else {
				// discovered, skip, if cost is greater
				if tentStartToCost >= startToCost[neighbor.LogicalP()] {
					continue
				}
				// esle update queue
				discovered.update(neighbor.LogicalP(), tentStartToCost+costEstimate(neighbor, to))
			}
			// we reach here if neighbour is just discovered or if we found a better path to it
			cameFrom[neighbor.LogicalP()] = current.node
			startToCost[neighbor.LogicalP()] = tentStartToCost
		}
	}

	return nil
}

func reconstructPath(cameFrom map[core.Point]core.Node, node core.Node) (path []core.Node) {
	fmt.Println("reconstructPath")
	for k, v := range cameFrom {
		fmt.Printf("%v came from %v\n", k, v.LogicalP())
	}
	path = append(path, node)
	for node, ok := cameFrom[node.LogicalP()]; ok; node, ok = cameFrom[node.LogicalP()] {
		// fmt.Printf("cur node: %v", node)
		path = append(path, node)
	}
	return
}
