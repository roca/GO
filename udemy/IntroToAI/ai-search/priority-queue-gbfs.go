package main

type PriorityQueueGBFS []*Node

func (pq PriorityQueueGBFS) Len() int {
	return len(pq)
}

func (pq PriorityQueueGBFS) Less(i, j int) bool {
	// GBFS uses only heuristic - lower Manhattan distance has higher priority
	return pq[i].CostToGoal < pq[j].CostToGoal
}

func (pq PriorityQueueGBFS) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueGBFS) Push(x any) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueueGBFS) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}
