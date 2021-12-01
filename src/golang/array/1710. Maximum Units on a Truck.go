package array

import "sort"

//Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
//Output: 8
func MaximumUnits(boxTypes [][]int, truckSize int) int{
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	var res int = 0
	for i := 0;i < len(boxTypes);i++{
		res += min_int(truckSize,boxTypes[i][0]) * boxTypes[i][1]
		truckSize -= boxTypes[i][0]
		if truckSize <= 0{
			break
		}
	}
	return res
}

//type sortBox [][]int
//
//func (s sortBox) Less(i, j int) bool {
//	return s[i][1] > s[j][1]
//}
//
//func (s sortBox) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//func (s sortBox) Len() int {
//	return len(s)
//}
//
//func MaximumUnits(boxTypes [][]int, truckSize int) int {
//	sort.Sort(sortBox(boxTypes))
//	var l int = len(boxTypes)
//	var res int = 0
//	for i := 0;i < l;i++{
//		if boxTypes[i][0] <= truckSize{
//			res += boxTypes[i][0] * boxTypes[i][1]
//			truckSize -= boxTypes[i][0]
//		}else{
//			res += truckSize * boxTypes[i][1]
//			truckSize = 0
//		}
//		if truckSize <= 0{
//			break
//		}
//	}
//	return res
//}