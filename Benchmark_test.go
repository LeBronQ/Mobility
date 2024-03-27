package Mobility

import (
	"math"
	"testing"
)

const (
	n        = 1000
	minSpeed = 10.0
	maxSpeed = 100.0
)

var arr []*Node

func GenerateNodes() {
	for i := 1; i <= n; i++ {
		pos := Nbox.RandomPosition3D()
		angle1 := RandomAngle()
		angle2 := RandomAngle()
		s := RandomSpeed(minSpeed, maxSpeed)
		var newSpeed Speed
		newSpeed.X = s * math.Cos(angle1) * math.Cos(angle2)
		newSpeed.Y = s * math.Cos(angle1) * math.Sin(angle2)
		newSpeed.Z = s * math.Sin(angle1)
		n := &Node{
			ID:    int64(i),
			Pos:   pos,
			Time:  0,
			V:     newSpeed,
			Model: "RandomWalk",
			Param: RandomWalkParam{
				MinSpeed: minSpeed,
				MaxSpeed: maxSpeed,
			},
		}
		arr = append(arr, n)
	}
}

// 单节点更新移动性模型
func BenchmarkSingleNode(b *testing.B) {
	n := &Node{
		ID: 1,
		Pos: Position{
			X: 200., Y: 200., Z: 200.,
		},
		Time: 10,
		V: Speed{
			X: 10., Y: 10., Z: 10.,
		},
		Model: "RandomWalk",
		Param: RandomWalkParam{
			MinSpeed: 1000,
			MaxSpeed: 5000,
		},
	}
	UpdatePosition(n)
}

func BenchmarkMultiNode(b *testing.B) {
	b.StopTimer()
	GenerateNodes()
	b.StartTimer()
	for _, n := range arr {
		UpdatePosition(n)
	}
}
