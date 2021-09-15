package best_meeting_point

func minTotalDistance(grid [][]int) int {
	rows := collectRows(grid)
	cols := collectCols(grid)
	return minDistance1D(rows) + minDistance1D(cols)
}

func minDistance1D(points []int) int {
	distance := 0
	i := 0
	j := len(points) - 1
	for i < j {
		distance += points[j] - points[i]
		i++
		j--
	}
	return distance
}

func collectRows(grid [][]int) []int {
	rows := make([]int, 0)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 1 {
				rows = append(rows, row)
			}
		}
	}
	return rows
}

func collectCols(grid [][]int) []int {
	cols := make([]int, 0)
	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < len(grid); row++ {
			if grid[row][col] == 1 {
				cols = append(cols, col)
			}
		}
	}
	return cols
}
