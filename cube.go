package hexgrid

type CubeCoord struct {
	X, Y, Z int
}

var _ Coord = &CubeCoord{}

func (c CubeCoord) ToOddq() (o OddqCoord) {
	o.Row = c.Z + (c.X-(c.X&1))/2
	o.Col = c.X
	return
}

func (c CubeCoord) ToCube() CubeCoord {
	return c
}
