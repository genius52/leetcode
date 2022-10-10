package array

//给你一个长度为 n 的 整数 数组 pref 。找出并返回满足下述条件且长度为 n 的数组 arr ：
//
//pref[i] = arr[0] ^ arr[1] ^ ... ^ arr[i].
func findArray(pref []int) []int {
	var l int = len(pref)
	var arr []int = make([]int,l)
	arr[0] = pref[0]
	var sum int = pref[0]
	for i := 1;i < l;i++{
		arr[i] = sum ^ pref[i]
		sum ^= arr[i]
	}
	return arr
}