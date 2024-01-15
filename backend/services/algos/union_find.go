package algos

type Subset struct {
	Rank   int
	Parent int
}

func Find(subs []Subset, root int) int {
	if subs[root].Parent != root {
		subs[root].Parent = Find(subs, subs[root].Parent)
	}
	return subs[root].Parent
}

func Union(subs []Subset, x, y int) {
	root1, root2 := Find(subs, x), Find(subs, y)
	rank1, rank2 := subs[root1].Rank, subs[root2].Rank

	if rank1 < rank2 {
		subs[root1].Parent = root2
	} else if rank1 > rank2 {
		subs[root2].Parent = root1
	} else {
		subs[root2].Parent = root1
		subs[root1].Rank++
	}
}
