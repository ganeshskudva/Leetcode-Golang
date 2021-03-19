package Medium 

func canVisitAllRooms(rooms [][]int) bool {
	if rooms == nil || len(rooms) == 0 {
		return false
	}

	mp := make(map[int]bool)
	check(rooms, 0, mp)

	return len(mp) == len(rooms)
}

func check(rooms [][]int, room int, mp map[int]bool) {
	if _, ok := mp[room]; ok {
		return
	}

	mp[room] = true

	for _, r := range rooms[room] {
		check(rooms, r, mp)
	}
}
