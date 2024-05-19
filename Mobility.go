package Mobility

import "fmt"

const (
	TimeSlot = 0.1
	TimeStep = 1.0
)

type Speed struct {
	X float64
	Y float64
	Z float64
}

type Position struct {
	X float64
	Y float64
	Z float64
}

type Node struct {
	ID    int64
	Pos   Position
	Time  uint64
	V     Speed
	Model string
	Param interface{}
}

var Nbox = NewBox(0, 4000, 0, 4000, 0, 4000)

func UpdatePosition(n Node) Node {
	if n.Time == 0 {
		switch n.Model {
		case "RandomWalk":
			n.V = InitRandomWalk(n.Param)
			n.Time = TimeStep / TimeSlot
		case "RandomWayPoint":
		case "GaussMarkov":
		default:
			fmt.Printf("Model does not exist")
		}
	}
	var newPos Position
	newPos.X = n.Pos.X + n.V.X*TimeSlot
	newPos.Y = n.Pos.Y + n.V.Y*TimeSlot
	newPos.Z = n.Pos.Z + n.V.Z*TimeSlot
	n.Pos, n.V = Nbox.BoundProcess(newPos, n.V, n.Pos)
	n.Time -= 1
	fmt.Printf("time:%d", n.Time)
	return n
}
