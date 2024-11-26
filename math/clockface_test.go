package clockface_test

import (
	"math"
	"testing"
	"time"

	clockface "github.com/nt2311-vn/TTD_GO/math"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, but want %v", got, want)
	}
}

func TestSecondsInRadiants(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadians(thirtySeconds)

	if want != got {
		t.Fatalf("Wanted %v radians, but got %v", want, got)
	}
}

func secondsInRadians(t time.Time) float64 {
	return 0
}
