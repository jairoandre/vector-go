package vector_go

import (
	"math"
)

type Vector2d struct {
	X float64
	Y float64
}

var (
	Zero = Vector2d{0, 0}
)

func NewVec2d(x, y float64) Vector2d {
	return Vector2d{
		X: x,
		Y: y,
	}
}

func NewVec2dFromInt(x, y int) Vector2d {
	return Vector2d{
		X: float64(x),
		Y: float64(y),
	}
}

func (v Vector2d) IntCoords() (int, int) {
	return int(v.X), int(v.Y)
}

func (v Vector2d) Add(o Vector2d) Vector2d {
	return Vector2d{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vector2d) Sub(o Vector2d) Vector2d {
	return Vector2d{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vector2d) Mul(scalar float64) Vector2d {
	return Vector2d{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vector2d) Div(scalar float64) Vector2d {
	return Vector2d{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vector2d) LenSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector2d) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

func (v Vector2d) SetMag(mag float64) Vector2d {
	return v.Normalize().Mul(mag)
}

func (v Vector2d) Normalize() Vector2d {
	vLen := v.Len()
	if vLen == 0 {
		return Zero
	}
	return Vector2d{
		X: v.X / vLen,
		Y: v.Y / vLen,
	}
}

func (v Vector2d) Rotate(theta float64) Vector2d {
	x := math.Cos(theta)*v.X - math.Sin(theta)*v.Y
	y := math.Sin(theta)*v.X + math.Cos(theta)*v.Y
	return Vector2d{x, y}
}

func (v Vector2d) PointsBetween(o Vector2d) []Vector2d {
	dx := v.X - o.X
	dy := v.Y - o.Y
	absDx := math.Abs(dx)
	absDy := math.Abs(dy)
	if dx == 0 && dy == 0 {
		return []Vector2d{}
	}
	xDiffIsLarger := absDx > absDy
	stepX := 1.0
	if dx < 0 {
		stepX = -1.0
	}
	stepY := 1.0
	if dy < 0 {
		stepY = -1.0
	}
	longerSideLength := math.Max(absDx, absDy)
	shorterSideLength := math.Min(absDx, absDy)
	slope := shorterSideLength / longerSideLength
	points := make([]Vector2d, 0)
	for i := 1.0; i <= longerSideLength; i++ {
		shorterSideIncrease := math.Round(i * slope)
		xIncrease := 0.0
		yIncrease := 0.0
		if xDiffIsLarger {
			xIncrease = i
			yIncrease = shorterSideIncrease
		} else {
			xIncrease = shorterSideIncrease
			yIncrease = i
		}
		toX := o.X + math.Round(xIncrease*stepX)
		toY := o.Y + math.Round(yIncrease*stepY)
		points = append(points, Vector2d{X: toX, Y: toY})
	}
	return points
}
