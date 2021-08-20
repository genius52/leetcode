package tree

import "math"

func pathInZigZagTree1(label int) []int{
	var res []int = []int{label}
	var layer_num int = math.Ilogb(float64(label))
	for layer_num > 0{
		var parent int = label/2
		layer_min := int(math.Pow(2,float64(layer_num - 1)))
		layer_max := int(math.Pow(2,float64(layer_num)) - 1)
		label = layer_max - (parent - layer_min)
		res = append([]int{label},res...)
		layer_num--
	}
	return res
}

func pathInZigZagTree(label int) []int {
	var res []int
	res = append(res, label)
	var layer_num int = math.Ilogb(float64(label))
	for layer_num > 1{
		var parent int = label/2
		layer_min := int(math.Pow(2,float64(layer_num - 1)))
		layer_max := int(math.Pow(2,float64(layer_num)) - 1)
		parent = layer_max - (parent - layer_min)
		label = parent
		layer_num--
		res = append(res, parent)
	}
	res = append(res, 1)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}