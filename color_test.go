// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"fmt"
	"testing"
)

func TestRGBToHSV(t *testing.T) {
	for r := 0; r < 256; r++ {
		for g := 0; g < 256; g++ {
			for b := 0; b < 256; b++ {
				hsvi := RGBToHSVi(uint8(r), uint8(g), uint8(b))
				r_, g_, b_ := hsvi.ToRGB()
				if r != int(r_) || g != int(g_) || b != int(b_) {
					panic(fmt.Errorf("%d, %d, %d", r, g, b))
				}
			}
		}
	}
	fmt.Println("factor", factor)
	fmt.Println("black", RGBToHSVi(0, 0, 0))
	fmt.Println("white", RGBToHSVi(255, 255, 255))
	fmt.Println("red", RGBToHSVi(255, 0, 0))
	fmt.Println("green", RGBToHSVi(0, 255, 0))
	fmt.Println("blue", RGBToHSVi(0, 0, 255))
	fmt.Println("max hue", RGBToHSVi(255, 0, 1))
}
