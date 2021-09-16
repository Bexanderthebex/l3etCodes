package graph_valid_tree

type Stack []int

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func validTree(n int, edges [][]int) bool {
	// For the graph to be a valid tree, it must have exactly n - 1 edges.
	// Any less, it can't possibly be fully connected
	// Any more, it has to contain cycles
	if len(edges) != n-1 {
		return false
	}

	nodesMap := make(map[int][]int)

	// construct the adjacency list
	for _, edge := range edges {
		vertex1 := edge[0]
		vertex2 := edge[1]
		nodesMap[vertex1] = append(nodesMap[vertex1], vertex2)
		nodesMap[vertex2] = append(nodesMap[vertex2], vertex1)
	}

	// set for the nodes that already have been visited
	visited := make(map[int]bool)
	stack := make(Stack, 0)
	stack.Push(0)
	visited[0] = true

	for !stack.IsEmpty() {
		node, _ := stack.Pop()
		for _, neighbor := range nodesMap[node] {
			if _, exists := visited[neighbor]; exists {
				continue
			}
			visited[neighbor] = true
			stack.Push(neighbor)
		}
	}

	return len(visited) == n
}
