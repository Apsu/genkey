/*
Copyright (C) 2021 Colin Hughes
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

var StaggerFlag bool
var SlideFlag bool
var DynamicFlag bool
var ImproveFlag bool
var ImproveLayout Layout

var FingerNames = [10]string{"LP", "LR", "LM", "LI", "RI", "RM", "RR", "RP", "LT", "RT"}

var Layouts map[string]Layout
var GeneratedFingermap map[Finger][]Pos
var GeneratedFingermatrix map[Pos]Finger

var SwapPossibilities []Pos

var Analyzed int

var Weight struct {
	FSpeed struct {
		SFB       float64
		DSFB      float64
		KeyTravel float64
		KPS       [10]float64
	}
	Dist struct {
		Lateral float64
	}
	Score struct {
		FSpeed       float64
		IndexBalance float64
		LSB          float64

		TrigramPrecision int
		LeftInwardRoll   float64
		LeftOutwardRoll  float64
		RightInwardRoll  float64
		RightOutwardRoll float64
		Alternate        float64
		Redirect         float64
		Onehand          float64
	}
}
