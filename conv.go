// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

func RGBFromBytes(r, g, b uint8) RGB {
	return RGB{float64(r) / 255, float64(g) / 255, float64(b) / 255}
}

func RGBFromHex(x uint32) RGB {
	return RGBFromBytes(
		uint8((x>>16)&0xFF),
		uint8((x>>8)&0xFF),
		uint8(x&0xFF))
}

func (cl *RGB) ToBytes() (r, g, b uint8) {
	return uint8(255*cl.R + 0.5), uint8(255*cl.G + 0.5), uint8(255*cl.B + 0.5)
}

func (cl *RGB) ToHSV() HSV {
	r, g, b := cl.R, cl.G, cl.B
	h, s, v := 0.0, 0.0, 0.0

	min := cl.Min()
	max := cl.Max()
	v, c := max, max-min
	if v == 0 { // black
		return HSV{H: -1, S: 0, V: 0}
	}
	s = c / v
	switch max {
	case r:
		h = (g - b) / c
		if h < 0 {
			h += 6
		}
	case g:
		h = 2 + (b-r)/c
	default: // b
		h = 4 + (r-g)/c
	}
	h /= 6 // 0..1
	return HSV{h, s, v}
}

func (cl *RGB) ToHCL() HCL {
	r, g, b := cl.R, cl.G, cl.B
	h, c, l := 0.0, 0.0, 0.0

	min := cl.Min()
	max := cl.Max()
	l, c = 0.5*(max+min), max-min
	if c == 0 {
		return HCL{-1, c, l}
	}
	switch max {
	case r:
		h = (g - b) / c
		if h < 0 {
			h += 6
		}
	case g:
		h = 2 + (b-r)/c
	default: // b
		h = 4 + (r-g)/c
	}
	h /= 6 // 0..1
	return HCL{h, c, l}
}

func (cl *HSV) ToRGB() RGB {
	h, s, v := cl.H, cl.S, cl.V
	r, g, b := float64(0), float64(0), float64(0)

	if s == 0 { // grey
		return RGB{v, v, v}
	}

	c := v * s
	min, max := v-c, v
	h *= 6 // sector 0 to 5
	off := c * (h - float64(int(h)))

	switch int(h) {
	case 0:
		r, g, b = max, min+off, min
	case 1:
		r, g, b = max-off, max, min
	case 2:
		r, g, b = min, max, min+off
	case 3:
		r, g, b = min, max-off, max
	case 4:
		r, g, b = min+off, min, max
	default: // case 5:
		r, g, b = max, min, max-off
	}
	return RGB{r, g, b}
}

func (cl *HCL) ToRGB() RGB {
	h, c, l := cl.H, cl.C, cl.L
	r, g, b := float64(0), float64(0), float64(0)

	if c == 0 {
		return RGB{l, l, l}
	}

	hc := 0.5 * c
	min, max := l-hc, l+hc
	h *= 6 // sector 0 to 5
	off := c * (h - float64(int(h)))

	switch int(h) {
	case 0:
		r, g, b = max, min+off, min
	case 1:
		r, g, b = max-off, max, min
	case 2:
		r, g, b = min, max, min+off
	case 3:
		r, g, b = min, max-off, max
	case 4:
		r, g, b = min+off, min, max
	default: // case 5:
		r, g, b = max, min, max-off
	}
	return RGB{r, g, b}

}

func (cl *RGB) Min() (min float64) {
	min = cl.R
	if cl.G < min {
		min = cl.G
	}
	if cl.B < min {
		min = cl.B
	}
	return min
}

func (cl *RGB) Max() (max float64) {
	max = cl.R
	if cl.G > max {
		max = cl.G
	}
	if cl.B > max {
		max = cl.B
	}
	return max
}

/*
func RGBToHSV(r_, g_, b_ uint8) HSV {
	r, g, b := float64(r_)/255, float64(g_)/255, float64(b_)/255
	h, s, b := float64(0), float64(0), float64(0)
	k := float64(0)

	if g < b {
		g, b = b, g
		k = -1.
	}

	if r < g {
		r, g = g, r
		k = -2./6. - k
	}

	chroma := r - min(g, b)
	h = abs(k+(g-b)/(6.*chroma+1e-20)) * 360
	s = chroma / (r + 1e-20)
	b = r
	return HSV{h, s, b}
}
*/
