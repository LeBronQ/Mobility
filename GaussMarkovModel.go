package Mobility

import (
	"github.com/gonum/stat/distuv"
	"math"
)

func GaussMarkov3D(iteration int64, variance float64, alpha float64, lSpeed float64, lastDir float64,
	lastPit float64, meanSpeed float64, meanDir float64, meanPit float64, lastSpeed Speed, pos Position) (Speed, Position, int64,
	float64, float64, float64, float64, float64, float64) {
	var newPos Position
	newPos.X = pos.X + lastSpeed.X*TimeStep
	newPos.Y = pos.Y + lastSpeed.Y*TimeStep
	newPos.Z = pos.Z + lastSpeed.Z*TimeStep
	newPos, _ = Nbox.BoundProcess(newPos, lastSpeed, pos)

	speedVar := distuv.Normal{Mu: 0, Sigma: variance * variance}.Rand()
	dirVar := distuv.Normal{Mu: 0, Sigma: variance * variance}.Rand()
	pitVar := distuv.Normal{Mu: 0, Sigma: variance * variance}.Rand()
	sqrtAlpha := math.Sqrt(1 - alpha*alpha)
	sn := alpha*lSpeed + (1-alpha)*meanSpeed + sqrtAlpha*speedVar
	dn := alpha*lastDir + (1-alpha)*meanDir + sqrtAlpha*dirVar
	pn := alpha*lastPit + (1-alpha)*meanPit + sqrtAlpha*pitVar

	var newSpeed Speed
	newSpeed.X = sn * math.Cos(dn) * math.Cos(pn)
	newSpeed.Y = sn * math.Sin(dn) * math.Cos(pn)
	newSpeed.Z = sn * math.Sin(pn)

	meanSpeed = (meanSpeed*(float64(iteration)) + sn) / (float64)(iteration+1)
	meanDir = (meanDir*(float64(iteration)) + dn) / (float64)(iteration+1)
	meanPit = (meanPit*(float64(iteration)) + pn) / (float64)(iteration+1)
	iteration += 1

	return newSpeed, newPos, iteration, sn, dn, pn, meanSpeed, meanDir, meanPit
}
