package allpathsleadstodestination

func leadsToDestinationIter(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	states := make([]int, n)
	stack := make([]int, 0, len(edges))
	stack = append(stack, source)
	for i := 0; i < n; i++ {
		states[i] = WHITE
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	for len(stack) > 0 {
		n := stack[len(stack)-1]

		if states[n] == GRAY {
			stack = stack[:len(stack)-1]
			states[n] = BLACK
			continue
		}

		if len(graph[n]) == 0 {
			if n == destination {
				stack = stack[:len(stack)-1]
				continue
			}

			return false
		}

		states[n] = GRAY

		for _, node := range graph[n] {
			// cycle detection. If we are going to visit it again while the parent is still in the stack, it is a cycle
			if states[node] == GRAY {
				return false
			}
			if states[node] == WHITE {
				stack = append(stack, node)
			}
		}

	}

	return true
}
