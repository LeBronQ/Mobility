package Mobility

import (
	"math"
	"math/rand"
	"time"
)

func RandomAngle() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() * 2 * math.Pi //
}

func RandomSpeed(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}
func calculateDistance3D(src Position, dst Position) float64 {
	dx := dst.X - src.X
	dy := dst.Y - src.Y
	dz := dst.Z - src.Z
	distance3D := math.Sqrt(dx*dx + dy*dy + dz*dz)
	return distance3D
}
