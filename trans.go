// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

func (cl *RGB) Reverse() {
	cl.R = 1 - cl.R
	cl.G = 1 - cl.G
	cl.B = 1 - cl.B
}

func (cl *RGB) Times(ratio float64) {
	if ratio > 1.0 {
		ratio = 1.0
	}
	cl.R *= ratio
	cl.G *= ratio
	cl.B *= ratio
}

func (cl *RGB) Plus(o RGB) RGB {
	return RGB{min(cl.R+o.R, 1.0), min(cl.G+o.G, 1.0), min(cl.B+o.B, 1.0)}
}

func (cl *HSV) Times(ratio float64) HSV {
	if ratio > 1.0 {
		ratio = 1.0
	}
	return HSV{cl.H * ratio, cl.S * ratio, cl.V * ratio}
}

func (cl *HSV) Plus(o HSV) HSV {
	h := cl.H + o.H
	if h > 360 {
		h -= 360
	}
	return HSV{h, min(cl.S+o.S, 1.0), min(cl.V+o.V, 1.0)}
}


