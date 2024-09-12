package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func BuildGraph(e int, edges [][]int, nodeAdjList map[int]*Node) *Node {
	if n, exists := nodeAdjList[e]; exists {
		return n
	}

	i := e - 1
	edge := edges[i]

	newNode := &Node{
		Val:       e,
		Neighbors: []*Node{},
	}

	nodeAdjList[e] = newNode

	for _, ne := range edge {
		newNeighbour := BuildGraph(ne, edges, nodeAdjList)
		newNode.Neighbors = append(newNode.Neighbors, newNeighbour)
	}

	return newNode
}
