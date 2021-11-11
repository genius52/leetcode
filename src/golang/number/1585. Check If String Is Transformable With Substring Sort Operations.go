package number

//输入：s = "84532", t = "34852"
//输出：true
//解释：你可以按以下操作将 s 转变为 t ：
//"84532" （从下标 2 到下标 3）-> "84352"
//"84352" （从下标 0 到下标 2） -> "34852"

//s = "84532", t = "34852"
//我们想把 s 变成 t，那么第一步我们需要把 s 的第一个字母变成 3，其实就可以从第一个字母遍历到是 3 的地方，
//如果这个过程中遇到比 3 小的字母，那么一定是不可能变过来的，因为怎么 sort，这个字母一定在 3 前面，所以就不行，
//如果没有的话，那么我们就标记一下 s 中 3 的下标，表明已经用过了，下次遍历的时候就跳过就好了。然后依次类推，
//每次都遍历一遍，时间复杂度为 O(n^2)O(n
func isTransformable(s string, t string) bool {
	var l1 int = len(s)
	var l2 int = len(t)
	if l1 != l2{
		return false
	}
	var record [10][]int
	for i := 0;i < l1;i++{
		record[s[i] - '0'] = append(record[s[i] - '0'],i)
	}
	for i := 0;i < l1;i++{
		num := int(t[i] - '0')
		if len(record[num]) == 0{
			return false
		}
		for j := 0;j < num;j++{
			//只需要判断最前面的就行
			if len(record[j]) > 0 && record[j][0] < record[num][0]{
				return false
			}
		}
		record[num] = record[num][1:]
	}
	return true
}