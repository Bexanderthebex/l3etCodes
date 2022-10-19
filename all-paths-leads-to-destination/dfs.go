package allpathsleadstodestination

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	states := make([]int, n)
	for i := 0; i < n; i++ {
		states[i] = WHITE
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	var leadsToDest func(currNode int, destination int) bool
	leadsToDest = func(currNode int, destination int) bool {

		if states[currNode] != WHITE {
			return states[currNode] == BLACK
		}

		// leaf case
		if len(graph[currNode]) == 0 {
			return currNode == destination
		}

		states[currNode] = GRAY

		for _, node := range graph[currNode] {
			if !leadsToDest(node, destination) {
				return false
			}
		}

		states[currNode] = BLACK
		return true
	}

	return leadsToDest(source, destination)
}
