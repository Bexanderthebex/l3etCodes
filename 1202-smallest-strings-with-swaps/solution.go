package smalleststringswithswaps

func smallestStringWithSwaps(s string, pairs [][]int) string {
	// process the pairs in the dsu
	string_len := len(s)
	dsu := NewDSU(string_len)
	for _, pair := range pairs {
			dsu.union(pair[0], pair[1])
	}
	
	// make an adjacency list to process the strings
	character_graph := make(map[int][]int)
	for i := 0; i < string_len; i++ {
			root := dsu.find(i)
			character_graph[root] = append(character_graph[root], i)        
	}
	
	// use the adjacency list to pick the characters in the string 
	answer := make([]string, string_len)
	for _, nodes := range character_graph {
			temp := make([]string, 0)
			for _, node := range nodes {
					temp = append(temp, string(s[node]))
			}
			
			sort.Strings(temp)
			
			sorted_string_index := 0
			for _, node := range nodes {
					answer[node] = temp[sorted_string_index]
					sorted_string_index++
			}
	}
	
	return strings.Join(answer[:], "")
}
