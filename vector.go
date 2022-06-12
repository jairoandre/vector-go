package vector_go

import "math"

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
