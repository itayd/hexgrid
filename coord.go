package hexgrid

// Based on https://www.redblobgames.com/grids/hexagons/.

type Coord interface {
	ToCube() CubeCoord
	ToOddq() OddqCoord
}

func AreEqual(a, b Coord) bool {
	switch aa := a.(type) {
	case OddqCoord:
		return aa == b.ToOddq()
	case CubeCoord:
		return aa == b.ToCube()
	default:
		return AreEqual(a.ToCube(), b)
	}
}

func AreNeighbours(a, b Coord) bool {
	return Distance(a, b) == 1
}

func Distance(a, b Coord) int {
	aa, bb := a.ToCube(), b.ToCube()

	return max3(
		abs(aa.X-bb.X),
		abs(aa.Y-bb.Y),
		abs(aa.Z-bb.Z),
	)
}

func Neighbours(c Coord) []Coord {
	ns := make([]Coord, 6)

	switch cc := c.(type) {
	case OddqCoord:
		deltas := [][]OddqCoord{
			/* even */ []OddqCoord{
				OddqCoord{Row: 0, Col: -1},
				OddqCoord{Row: 1, Col: -1},
				OddqCoord{Row: 1, Col: 0},
				OddqCoord{Row: 0, Col: 1},
				OddqCoord{Row: -1, Col: 0},
				OddqCoord{Row: -1, Col: -1},
			},
			/* odd */ []OddqCoord{
				OddqCoord{Row: 0, Col: -1},
				OddqCoord{Row: 1, Col: 0},
				OddqCoord{Row: 1, Col: 1},
				OddqCoord{Row: 0, Col: 1},
				OddqCoord{Row: -1, Col: 1},
				OddqCoord{Row: -1, Col: 0},
			},
		}

		for i, delta := range deltas[cc.Col&1] {
			ns[i] = OddqCoord{
				Row: cc.Row + delta.Row,
				Col: cc.Col + delta.Col,
			}
		}

	case CubeCoord:
		deltas := []CubeCoord{
			CubeCoord{X: 1, Y: -1},
			CubeCoord{X: 1, Z: -1},
			CubeCoord{Y: 1, Z: -1},
			CubeCoord{X: -1, Y: 1},
			CubeCoord{X: -1, Z: 1},
			CubeCoord{Y: -1, Z: 1},
		}

		for i, delta := range deltas {
			ns[i] = CubeCoord{
				X: cc.X + delta.X,
				Y: cc.Y + delta.Y,
				Z: cc.Z + delta.Z,
			}
		}

	default:
		return Neighbours(c.ToCube())
	}

	return ns
}
