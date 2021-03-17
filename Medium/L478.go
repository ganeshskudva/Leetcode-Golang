package Medium 

type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius:  radius,
		xCenter: x_center,
		yCenter: y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	length := math.Sqrt(rand.Float64()) * this.radius
	deg := rand.Float64() * 2 * math.Pi
	x := this.xCenter + length*math.Cos(deg)
	y := this.yCenter + length*math.Sin(deg)

	return []float64{x, y}
}
