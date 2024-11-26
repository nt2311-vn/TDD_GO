package clockface

import "time"

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 240}
}

func secondsInRadians(t time.Time) float64 {
	return 0
}
