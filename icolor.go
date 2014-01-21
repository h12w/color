// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

/*
Translaction from:
http://code.google.com/p/streumix-frei0r-goodies/wiki/Integer_based_RGB_HSV_conversion
*/

import (
	"fmt"
)

const (
	vshift = 4
	factor = 255 << vshift
	k3     = 1 << (vshift - 1)

	Fac = factor
)

type HSVi struct {
	H int16
	S int16
	V int16
}

func RGBToHSVi(r_, g_, b_ uint8) HSVi {
	r, g, b := int(r_), int(g_), int(b_)
	var min, max int
	if r > g {
		if r > b {
			max = r
			if g > b {
				min = b
			} else {
				min = g
			}
		} else {
			max = b
			min = g
		}
	} else { // r < g
		if g > b {
			max = g
			if r > b {
				min = b
			} else {
				min = r
			}
		} else {
			max = b
			min = r
		}
	}
	chroma := max - min

	v := max << vshift

	var s int
	if max != 0 {
		s = factor * chroma / max
	} else {
		return HSVi{}
	}

	var h int
	chroma6 := 6 * chroma
	switch max {
	case min:
		h = 0
	case r:
		h = factor * (chroma6 + g - b) / chroma6
		if h >= factor {
			h -= factor
		}
	case g:
		h = factor * (2*chroma + b - r) / chroma6
	case b:
		h = factor * (4*chroma + r - g) / chroma6
	}
	return HSVi{int16(h), int16(s), int16(v)}
}

func (hsv *HSVi) ToRGB() (r_, g_, b_ uint8) {
	var r, g, b int
	h, s, v := int(hsv.H), int(hsv.S), int(hsv.V)

	if s == 0 {
		vv := uint8(v >> vshift)
		return vv, vv, vv
	}

	i := 6 * h
	ih := i / factor
	m := v * (factor - s) / factor >> vshift
	f := i - (i/factor)*factor // factorial part * factor
	var x int
	if (ih & 1) == 1 { // ih = 1,3,5
		x = (v*(factor*factor-s*f)/(factor*factor) + k3) >> vshift

	} else {
		x = (v*(factor*factor-s*(factor-f))/(factor*factor) + k3) >> vshift
	}
	v = v >> vshift

	switch ih {
	case 0:
		r, g, b = v, x, m
	case 1:
		r, g, b = x, v, m
	case 2:
		r, g, b = m, v, x
	case 3:
		r, g, b = m, x, v
	case 4:
		r, g, b = x, m, v
	case 5:
		r, g, b = v, m, x
	}
	return uint8(r), uint8(g), uint8(b)
}

func absi(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func p(v ...interface{}) {
	fmt.Println(v...)
}
