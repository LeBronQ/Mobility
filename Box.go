package Mobility

import (
	"math"
	"math/rand"
	"time"
)

type Box struct {
	xMin float64
	xMax float64
	yMin float64
	yMax float64
	zMin float64
	zMax float64
}

func NewBox(_xMin float64, _xMax float64, _yMin float64, _yMax float64, _zMin float64, _zMax float64) *Box {
	box := new(Box)
	box.xMin = _xMin
	box.xMax = _xMax
	box.yMin = _yMin
	box.yMax = _yMax
	box.zMin = _zMin
	box.zMax = _zMax
	return box
}

func (box *Box) xIsInside(pos Position) bool {
	return pos.X <= box.xMax && pos.X >= box.xMin
}
func (box *Box) yIsInside(pos Position) bool {
	return pos.Y <= box.yMax && pos.Y >= box.yMin
}
func (box *Box) zIsInside(pos Position) bool {
	return pos.Z <= box.zMax && pos.Z >= box.zMin
}

func (box *Box) BoundCalculation(v float64, init float64, dir string) float64 {
	minBound, maxBound := 0.0, 0.0
	if dir == "x" {
		minBound = box.xMin
		maxBound = box.xMax
	} else if dir == "y" {
		minBound = box.yMin
		maxBound = box.yMax
	} else {
		minBound = box.zMin
		maxBound = box.zMax
	}
	delta := math.Abs(v * TimeSlot)
	a := maxBound - minBound
	if v > 0 {
		n := math.Floor(delta / (2 * a))
		left := delta - 2*a*n
		if left < maxBound-init {
			return left + init
		} else if left < maxBound-init+a {
			return maxBound + maxBound - init - left
		} else {
			return left - a - (maxBound - init) + minBound
		}
	} else {
		n := math.Floor(delta / (2 * a))
		leftX := delta - 2*a*n
		if leftX < init-minBound {
			return init - leftX
		} else if leftX < init-minBound+a {
			return minBound + leftX - (init - minBound)
		} else {
			return maxBound - (leftX - (init - minBound) - a)
		}
	}
}

func (box *Box) BoundProcess(newPos Position, speed Speed, pos Position) (Position, Speed) {
	if !box.xIsInside(newPos) {
		newPos.X = box.BoundCalculation(speed.X, pos.X, "x")
		speed.X = -speed.X
	}
	if !box.yIsInside(newPos) {
		newPos.Y = box.BoundCalculation(speed.Y, pos.Y, "y")
		speed.Y = -speed.Y
	}
	if !box.zIsInside(newPos) {
		newPos.Z = box.BoundCalculation(speed.Z, pos.Z, "z")
		speed.Z = -speed.Z
	}
	return newPos, speed
}

func (box *Box) RandomPosition3D() Position {
	rand.Seed(time.Now().UnixNano())
	var newPosition Position
	newPosition.X = box.xMin + rand.Float64()*(box.xMax-box.xMin)
	newPosition.Y = box.yMin + rand.Float64()*(box.yMax-box.yMin)
	newPosition.Z = box.zMin + rand.Float64()*(box.zMax-box.zMin)
	return newPosition
}
