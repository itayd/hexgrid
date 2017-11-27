package hexgrid

type OddqCoord struct {
	Row, Col int
}

var _ Coord = &OddqCoord{}

func (c OddqCoord) ToOddq() OddqCoord {
	return c
}

func (c OddqCoord) ToCube() (q CubeCoord) {
	q.X = c.Col
	q.Z = c.Row - (c.Col-(c.Col&1))/2
	q.Y = -q.X - q.Z
	return
}
