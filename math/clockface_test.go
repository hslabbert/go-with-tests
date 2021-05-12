package clockface

import (
	"math"
	"testing"
	"time"
)

/*
func TestSecondHand(t *testing.T) {
	cases := []struct {
		Name  string
		Time  time.Time
		Point Point
	}{
		{
			Name:  "At midnight",
			Time:  time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC),
			Point: Point{X: 150, Y: 150 - 90},
		},
		{
			Name:  "At 30 seconds",
			Time:  time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC),
			Point: Point{X: 150, Y: 150 + 90},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			want := test.Point
			got := SecondHand(test.Time)

			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}
*/

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadius(thirtySeconds)

	if want != got {
		t.Fatalf("wanted %v radians, but got %v", got, want)
	}

}
