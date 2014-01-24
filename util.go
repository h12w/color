// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"math"
)

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func min3(fs ...float64) float64 {
	if len(fs) == 0 {
		return 0
	}
	min := float64(math.MaxFloat32)
	for _, f := range fs {
		if f < min {
			min = f
		}
	}
	return min
}

func max3(fs ...float64) float64 {
	if len(fs) == 0 {
		return 0
	}
	max := float64(-math.MaxFloat32)
	for _, f := range fs {
		if f > max {
			max = f
		}
	}
	return max
}

// Algorithm from here:
// http://gamedev.stackexchange.com/questions/46463/is-there-an-optimum-set-of-colors-for-10-players
func DistinguishableColor(i int, s, b float64) HSV {
	hue := math.Mod(360*0.618033988749895*float64(i), 360.0)
	return HSV{hue, s, b}
}
