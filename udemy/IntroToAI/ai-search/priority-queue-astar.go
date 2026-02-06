package main

type PriorityQueueAstar []*Node

func (pq PriorityQueueAstar) Len() int {
	return len(pq)
}

func (pq PriorityQueueAstar) Less(i, j int) bool {
	return int(pq[i].EstimatedCostToGoal) < int(pq[j].EstimatedCostToGoal)
}

func (pq PriorityQueueAstar) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueAstar) Push(x any) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueueAstar) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}
