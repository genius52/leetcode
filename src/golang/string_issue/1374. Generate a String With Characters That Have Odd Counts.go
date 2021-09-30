package string_issue

func generateTheString(n int) string {
	var res string
	if n % 2 == 1{
		for i := 0;i < n;i++{
			res += "a"
		}
		return res
	}else{
		for i := 0;i < n - 1;i++{
			res += "a"
		}
		res += "b"
		return res
	}
}