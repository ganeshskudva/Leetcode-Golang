package Medium

func minimumFuelCost(roads [][]int, seats int) int64 {
	res, mp := int64(0), make(map[int][]int)

	for _, r := range roads {
		mp[r[0]] = append(mp[r[0]], r[1])
		mp[r[1]] = append(mp[r[1]], r[0])
	}

	solve(0, -1, int64(seats), mp, &res)
	return res
}

func solve(node, parent int, seats int64, mp map[int][]int, res *int64) int64 {
	rep := int64(1)

	if _, ok := mp[node]; !ok {
		return rep
	}

	for _, n := range mp[node] {
		if n != parent {
			rep += solve(n, node, seats, mp, res)
		}
	}

	if node != 0 {
		// math.ceil for int64
		*res += (rep + seats - 1) / seats
	}

	return rep
}
