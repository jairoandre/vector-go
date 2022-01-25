package vector_go

type Vector2d struct {
	X float64
	Y float64
}

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
