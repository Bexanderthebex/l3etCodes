package mincosttoconnectallpoints

type dsu struct {
	group []int
	rank  []int
	total int
}

func (d *dsu) find(node int) int {
	if d.group[node] == node {
		return node
	}
	d.group[node] = d.find(d.group[node])
	return d.group[node]
}

func (d *dsu) union(x, y int) {
	rootX := d.find(x)
	rootY := d.find(y)
	if rootX != rootY {
		if d.rank[rootX] > d.rank[rootY] {
			d.group[rootY] = rootX
		} else if d.rank[rootX] < d.rank[rootY] {
			d.group[rootX] = rootY
		} else {
			d.group[rootY] = rootX
			d.rank[rootX] += 1
		}

		d.total -= 1
	}
}

func (d *dsu) connected(x, y int) bool {
	return d.find(x) == d.find(y)
}

func (d *dsu) getUniqueRoots() int {
	return d.total
}

func newDsu(size int) dsu {
	d := dsu{group: make([]int, size), rank: make([]int, size), total: size}
	for i := range d.group {
		d.group[i] = i
		d.rank[i] = 1
	}
	return d
}

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 || isConnected == nil {
		return 0
	}
	dset := newDsu(len(isConnected))
	for row := range isConnected {
		for col := row + 1; col < len(isConnected); col++ {
			if isConnected[row][col] == 1 {
				dset.union(row, col)
			}
		}
	}
	return dset.getUniqueRoots()
}
