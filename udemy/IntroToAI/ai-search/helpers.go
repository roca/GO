package main

func inExplored(needle Point, haystack []Point) bool {
	for _, point := range haystack {
		if point.Row == needle.Row && point.Col == needle.Col {
			return true
		}
	}
	return false
}
