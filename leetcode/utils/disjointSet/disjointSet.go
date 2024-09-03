package disjointset

type Union struct {
	Root []int
	Rank []int
}

func NewUnion(n int) *Union {
	root := make([]int, n)
	for i := range root {
		root[i] = i
	}

	return &Union{
		Root: root,
		Rank: make([]int, n),
	}
}

func (u *Union) Find(i int) int {
	if i == u.Root[i] {
		return i
	}

	return u.Find(u.Root[i])
}

func (u *Union) Union(i, j int) {
	root_i := u.Find(i)
	root_j := u.Find(j)

	if root_i != root_j {
		switch {
		case u.Rank[root_i] > u.Rank[root_j]:
			u.Root[root_j] = root_i
		case u.Rank[root_i] < u.Rank[root_j]:
			u.Root[root_i] = root_j
		case u.Rank[root_i] == u.Rank[root_j]:
			u.Root[root_j] = root_i
			u.Rank[root_i]++
		}
	}
}

func (u *Union) Connected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}
