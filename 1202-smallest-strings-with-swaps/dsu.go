package smalleststringswithswaps

type dsu struct {
	root []int
	height []int
}

func (d *dsu) find(node int) int {
	if d.root[node] == node {
		return node
	}
	d.root[node] = d.find(d.root[node])
	return d.root[node]
}

func (d *dsu) union(x, y int) {
	rootX := d.find(x)
	rootY := d.find(y)
	if rootX != rootY {
		if d.height[rootX] > d.height[rootY] {
			d.root[rootY] = rootX 
		} else if d.height[rootY] < d.height[rootX] {
			d.root[rootX] = rootY
		} else {
			d.root[rootY] = rootX
			d.height[rootX] += 1
		}
	}
}

func (d *dsu) connected(x, y int) bool {
	return d.find(x) == d.find(y)
}

func NewDSU(numberOfNodes int) dsu {
	d := dsu{
		root: make([]int, numberOfNodes),
		height: make([]int, numberOfNodes)
	}

	for i := 0; i < numberOfNodes; i++ {
		d.root[i] = i
		d.height[i] = 1
	}

	return d
}
