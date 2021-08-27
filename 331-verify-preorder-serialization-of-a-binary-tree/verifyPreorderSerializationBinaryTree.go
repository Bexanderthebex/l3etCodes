package leetcode_331

func isValidSerialization(preorder string) bool {
	slots := 1

	for i, s := range preorder {
		if s == ',' {
			slots--

			if slots < 0 {
				return false
			}

			if preorder[i-1] != '#' {
				slots += 2
			}
		}
	}

	if preorder[len(preorder)-1] == '#' {
		slots--
	} else {
		slots++
	}
	return slots == 0
}
