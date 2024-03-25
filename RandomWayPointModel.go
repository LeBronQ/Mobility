package Mobility

import (
	"math/rand"
)

func RandomWayPoint3D(pos Position, minSpeed float64, maxSpeed float64, StayTimeLimit float64) (Speed, Position, float64, float64) {
	dst := Nbox.RandomPosition3D()
	s := RandomSpeed(minSpeed, maxSpeed)
	travelingTime := CalculateDistance3D(pos, dst) / s
	var newSpeed Speed
	newSpeed.X = (dst.X - pos.X) / travelingTime
	newSpeed.Y = (dst.Y - pos.Y) / travelingTime
	newSpeed.Z = (dst.Z - pos.Z) / travelingTime
	Stay := rand.Float64() * StayTimeLimit
	return newSpeed, dst, travelingTime, Stay
}
