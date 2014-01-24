// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
r - red
g - green
b - blue
h - hue
s - satuation
v - brightness (value)
c - chroma
l - lightness
*/
package color

type RGB struct {
	R, G, B float64
}

type HSV struct {
	H, S, V float64
}

type HCL struct {
	H, C, L float64
}
