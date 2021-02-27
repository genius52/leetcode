package number

//Given a positive integer, return its corresponding column title as appear in an Excel sheet.
//
//For example:
//
//    1 -> A
//    2 -> B
//    3 -> C
//    ...

//    26 -> Z
//    27 -> AA
//    28 -> AB
//    ...
//Example 1:
//
//Input: 1
//Output: "A"
//Example 2:
//
//Input: 28
//Output: "AB"
//Example 3:
//
//Input: 701
//Output: "ZY"
func ConvertToTitle(n int) string {
	var dict map[int]string = make(map[int]string)
	for i := 0;i < 26;i++{
		dict[i] = string('A' + i)
	}
	var s string
	for n > 0{
		//cur := (n - 1) % 26 + 'A'
		//s = string(cur) + s
		cur := (n - 1) % 26
		s = dict[cur] + s
		n = (n - 1) / 26
	}
	return s
}