package string_issue

func largestGoodInteger(num string) string {
	var l int = len(num)
	var res string
	for i := 1;i < l - 1;i++{
		if num[i] == num[i - 1] && num[i] == num[i + 1]{
			if num[i - 1:i + 2] > res{
				res = num[i - 1:i + 2]
			}
		}
	}
	return res
}