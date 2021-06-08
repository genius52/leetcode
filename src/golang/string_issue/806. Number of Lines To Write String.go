package string_issue

//Input: widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10], s = "bbbcccdddaaa"
//Output: [2,4]
//Explanation: You can write s as follows:
//bbbcccdddaa  // 98 pixels wide
//a            // 4 pixels wide
//There are a total of 2 lines, and the last line is 4 pixels wide.
func numberOfLines(widths []int, s string) []int {
	var res []int = make([]int,2)
	var lines int = 1
	var offset int = 0
	for _,c := range s{
		cur_width := widths[c - 'a']
		if offset + cur_width > 100{
			lines++
			offset = cur_width
		}else{
			offset += cur_width
		}
	}
	res[0] = lines
	res[1] = offset
	return res
}