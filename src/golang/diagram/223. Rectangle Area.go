package diagram

import "math"

func ComputeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	var left_edge int = int(math.Max(float64(A),float64(E)))
	var right_edge int = int(math.Min(float64(C),float64(G)))
	var top_edge int = int(math.Min(float64(D),float64(H)))
	var bottom_edge int = int(math.Max(float64(B),float64(F)))
	var width int = 0
	if(left_edge < right_edge){
		width = right_edge - left_edge
	}
	var height int = 0
	if(top_edge > bottom_edge){
		height = top_edge - bottom_edge
	}
	return (C-A)*(D-B) + (G-E)*(H-F) - width*height
}
