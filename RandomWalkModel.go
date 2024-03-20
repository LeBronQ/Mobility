package Mobility

import (
	"math"
)

type RandomWalkParam struct {
	MinSpeed float64
	MaxSpeed float64
}

func RandomWalk3D(p RandomWalkParam) Speed {
	angle1 := RandomAngle()
	angle2 := RandomAngle()
	s := RandomSpeed(p.MinSpeed, p.MaxSpeed)
	var newSpeed Speed
	newSpeed.X = s * math.Cos(angle1) * math.Cos(angle2)
	newSpeed.Y = s * math.Cos(angle1) * math.Sin(angle2)
	newSpeed.Z = s * math.Sin(angle1)
	return newSpeed
}

func InitRandomWalk(Param interface{}) Speed {
	var p, ok = Param.(RandomWalkParam)
	if Param == nil || !ok {
		p.MinSpeed = 0
		p.MaxSpeed = 10
	}
	sp := RandomWalk3D(p)
	return sp
}
