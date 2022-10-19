package determinetotalprovinces

type dsu struct {
	group []int
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
		d.group[rootY] = rootX
	}
}

func (d *dsu) connected(x, y int) bool {
	return d.find(x) == d.find(y)
}

func (d *dsu) getUniqueRoots() int {
	total := 0
	alreadyCounted := make(map[int]struct{}, len(d.group))
	for _, n := range d.group {
		root := d.find(n)
		if _, counted := alreadyCounted[root]; counted {
			continue
		}

		alreadyCounted[root] = struct{}{}
		total += 1
	}
	return total
}

func newDsu(size int) dsu {
	d := dsu{group: make([]int, size)}
	for i := range d.group {
		d.group[i] = i
	}
	return d
}

func findCircleNum(isConnected [][]int) int {
	dset := newDsu(len(isConnected))
	for i := range isConnected {
		for j := range isConnected[i] {
			if isConnected[i][j] != 1 || i == j {
				continue
			}
			if dset.connected(i, j) {
				continue
			}
			dset.union(i, j)
		}
	}
	return dset.getUniqueRoots()
}
