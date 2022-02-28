package string_issue

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	var l int = len(s)
	var hash int = 0
	var p int = 1
	for i := l - 1;i > l - 1 - k;i--{
		hash += int(s[i] - 'a' + 1) * p
		hash %= modulo
		p *= power
	}
	for i := l - 1 - k;i > k;i--{

	}
	return s
}