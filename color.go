// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"math"
)

type RGBf struct {
	R, G, B float64
}

func RGBfFromBytes(r, g, b uint8) RGBf {
	return RGBf{float64(r) / 255, float64(g) / 255, float64(b) / 255}
}

func RGBfFromHex(x uint32) RGBf {
	return RGBfFromBytes(
		uint8((x>>16)&0xFF),
		uint8((x>>8)&0xFF),
		uint8(x&0xFF))
}

func (c *RGBf) ToBytes() (r, g, b uint8) {
	return uint8(255 * c.R), uint8(255 * c.G), uint8(255 * c.B)
}

func (c *RGBf) Times(ratio float64) RGBf {
	if ratio > 1.0 {
		ratio = 1.0
	}
	return RGBf{c.R * ratio, c.G * ratio, c.B * ratio}
}

func (c *RGBf) Min() (min float64) {
	min = c.R
	if c.G < min {
		min = c.G
	}
	if c.B < min {
		min = c.B
	}
	return min
}

func (c *RGBf) Max() (max float64) {
	max = c.R
	if c.G > max {
		max = c.G
	}
	if c.B > max {
		max = c.B
	}
	return max
}

func (c *RGBf) Plus(o RGBf) RGBf {
	return RGBf{min(c.R+o.R, 1.0), min(c.G+o.G, 1.0), min(c.B+o.B, 1.0)}
}

func (c *RGBf) ToHSVf() HSVf {
	return RGBfToHSVf(c.ToBytes())
}

func RGBfToHSVf(r_, g_, b_ uint8) HSVf {
	r, g, b := float64(r_)/255, float64(g_)/255, float64(b_)/255
	h, s, v := float64(0), float64(0), float64(0)
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
	v = r
	return HSVf{h, s, v}
}

func (c *RGBf) ToHSVfSlow() HSVf {
	r, g, b := c.R, c.G, c.B
	h, s, v := 0.0, 0.0, 0.0 // v

	min := c.Min()
	max := c.Max()
	delta := max - min

	v = max
	if max != 0 {
		s = delta / max // s
	} else {
		// r = g = b = 0        // s = 0, v is undefined
		s = 0
		h = -1
		return HSVf{h, s, v}
	}

	if r == max {
		h = (g - b) / delta // between yellow & magenta
	} else if g == max {
		h = 2 + (b-r)/delta // between cyan & yellow
	} else {
		h = 4 + (r-g)/delta // between magenta & cyan
	}
	h *= 60 // degrees
	if h < 0 {
		h += 360
	}
	return HSVf{h, s, v}
}

type HSVf struct {
	H, S, V float64
}

func (c *HSVf) Times(ratio float64) HSVf {
	if ratio > 1.0 {
		ratio = 1.0
	}
	return HSVf{c.H * ratio, c.S * ratio, c.V * ratio}
}

func (c *HSVf) Plus(o HSVf) HSVf {
	h := c.H + o.H
	if h > 360 {
		h -= 360
	}
	return HSVf{h, min(c.S+o.S, 1.0), min(c.V+o.V, 1.0)}
}

func (c *HSVf) ToRGBfBytes() (r_, g_, b_ uint8) {
	h, s, v := c.H, c.S, c.V
	r, g, b := float64(0), float64(0), float64(0)

	if s == 0 {
		// achromatic (grey)
		v_ := uint8(v*255 + 0.5)
		return v_, v_, v_
	}

	h /= 60 // sector 0 to 5
	i := float64(int(h))
	f := h - i // factorial part of h
	p := v * (1 - s)
	q := v * (1 - s*f)
	t := v * (1 - s*(1-f))

	switch i {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	default: // case 5:
		r, g, b = v, p, q
	}
	return uint8(r*255 + 0.5), uint8(g*255 + 0.5), uint8(b*255 + 0.5)
}

func (c *HSVf) ToRGBf() RGBf {
	h, s, v := c.H, c.S, c.V
	r, g, b := float64(0), float64(0), float64(0)

	if s == 0 {
		// achromatic (grey)
		return RGBf{v, v, v}
	}

	h /= 60 // sector 0 to 5
	i := float64(int(h))
	f := h - i // factorial part of h
	p := v * (1 - s)
	q := v * (1 - s*f)
	t := v * (1 - s*(1-f))

	switch i {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	default: // case 5:
		r, g, b = v, p, q
	}
	return RGBf{r, g, b}
}

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
func DistinguishableColor(i int, s, v float64) HSVf {
	hue := math.Mod(360*0.618033988749895*float64(i), 360.0)
	return HSVf{hue, s, v}
}
