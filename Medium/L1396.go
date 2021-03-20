package Medium

type StationTime struct {
	station string
	time    int
}

type UndergroundSystem struct {
	ids map[int]StationTime
	avg map[[2]string][2]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{map[int]StationTime{}, map[[2]string][2]int{}}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.ids[id] = StationTime{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	start := this.ids[id]
	delete(this.ids, id)
	key := [2]string{start.station, stationName}
	v := this.avg[key]
	this.avg[key] = [2]int{v[0] + t - start.time, v[1] + 1}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	v := this.avg[[2]string{startStation, endStation}]
	return float64(v[0]) / float64(v[1])
}
