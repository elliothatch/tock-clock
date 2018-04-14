package main

import "github.com/lucasb-eyer/go-colorful"
//import "image"
//import "image/draw"
//import "image/png"
//import "os"
//import "strconv"

// This table contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]
type GradientTable []struct {
    Col colorful.Color
    Pos float64
}

// This is the meat of the gradient computation. It returns a HCL-blend between
// the two colors around `t`.
// Note: It relies heavily on the fact that the gradient keypoints are sorted.
func (self GradientTable) GetInterpolatedColorFor(t float64) colorful.Color {
    for i := 0 ; i < len(self) - 1 ; i++ {
        c1 := self[i]
        c2 := self[i+1]
        if c1.Pos <= t && t <= c2.Pos {
            // We are in between c1 and c2. Go blend them!
            t := (t - c1.Pos)/(c2.Pos - c1.Pos)
            return c1.Col.BlendHcl(c2.Col, t).Clamped()
        }
    }

    // Nothing found? Means we're at (or past) the last gradient keypoint.
    return self[len(self)-1].Col
}

// This is a very nice thing Golang forces you to do!
// It is necessary so that we can write out the literal of the colortable below.
func MustParseHex(s string) colorful.Color {
    c, err := colorful.Hex(s)
    if err != nil {
        panic("MustParseHex: " + err.Error())
    }
    return c
}

//package main;
//import (
	//"math"
//)

//// values are floats between 0 and 1
//type HsvColor struct {
	//h float32;
	//s float32;
	//v float32;
//}

//// values are floats between 0 and 1
//type RgbColor struct {
	//r float32;
	//g float32;
	//b float32;
//}

//func HsvToRgb(color HsvColor) RgbColor {
	//h, s, v := float64(color.h), float64(color.s), float64(color.v);
	//r, g, b := 0.0, 0.0, 0.0;

	//i := math.Floor(h * 6.0);
	//f := h * 6 - i;
	//p := v * (1 - s);
	//q := v * (1 - f * s);
	//t := v * (1 - (1 - f) * s);

	//switch int(i)%6 {
		//case 0:
			//r = v;
			//g = t;
			//b = p;

		//case 1:
			//r = q;
			//g = v;
			//b = p;

		//case 2:
			//r = p;
			//g = v;
			//b = t;

		//case 3:
			//r = p;
			//g = q;
			//b = v;

		//case 4:
			//r = t;
			//g = p;
			//b = v;

		//case 5:
			//r = v;
			//g = p;
			//b = q;
	//}

	//return RgbColor{float32(r), float32(g), float32(b)};
//}

