package shape

import "math"

/** 
* Created by wanjx in 2019/4/14 21:41
**/

type Point struct {
	X int
	Y int
}

func (p *Point) Distance(q *Point) int {
	return int(math.Sqrt(math.Pow(float64(q.X - p.X), 2.0) + math.Pow(float64(q.Y - p.Y), 2.0)))
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{X: p.X - q.X, Y: p.Y - q.Y}
}


type Path []Point

func (p Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range p {
		p[i] = op(p[i], offset)
	}

}