// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"fmt"
	"testing"
)

func testConv(conv func(r, g, b uint8) (r_, g_, b_ uint8)) {
	for ir := 0; ir < 256; ir++ {
		for ig := 0; ig < 256; ig++ {
			for ib := 0; ib < 256; ib++ {
				r, g, b := uint8(ir), uint8(ig), uint8(ib)
				r_, g_, b_ := conv(r, g, b)
				if r != r_ || g != g_ || b != b_ {
					panic(fmt.Errorf("(%d, %d, %d) -> (%d, %d, %d)", r, g, b, r_, g_, b_))
				}
			}
		}
	}
}

func TestHSVf(t *testing.T) {
	testConv(func(r, g, b uint8) (r_, g_, b_ uint8) {
		rgb := RGBFromBytes(r, g, b)
		hsv := rgb.ToHSV()
		rgb_ := hsv.ToRGB()
		return rgb_.ToBytes()
	})
}

func TestHCL(t *testing.T) {
	testConv(func(r, g, b uint8) (r_, g_, b_ uint8) {
		rgb := RGBFromBytes(r, g, b)
		hcl := rgb.ToHCL()
		rgb_ := hcl.ToRGB()
		return rgb_.ToBytes()
	})
}

func TestRGBToHSV(t *testing.T) {
	testConv(func(r, g, b uint8) (r_, g_, b_ uint8) {
		hsvi := RGBToHSVi(r, g, b)
		return hsvi.ToRGB()
	})
	fmt.Println("factor", factor)
	fmt.Println("black", RGBToHSVi(0, 0, 0))
	fmt.Println("white", RGBToHSVi(255, 255, 255))
	fmt.Println("red", RGBToHSVi(255, 0, 0))
	fmt.Println("green", RGBToHSVi(0, 255, 0))
	fmt.Println("blue", RGBToHSVi(0, 0, 255))
	fmt.Println("max hue", RGBToHSVi(255, 0, 1))
}
