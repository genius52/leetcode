package string_issue

//输入：s = "))()))", locked = "010100"
//输出：true
//解释：locked[1] == '1' 和 locked[3] == '1' ，所以我们无法改变 s[1] 或者 s[3] 。
//我们可以将 s[0] 和 s[4] 变为 '(' ，不改变 s[2] 和 s[5] ，使 s 变为有效字符串。
//If locked[i] is '1', you cannot change s[i].
//But if locked[i] is '0', you can change s[i] to either '(' or ')'.
func CanBeValid(s string, locked string) bool {
	var l int = len(s)
	if l|1 == l {
		return false
	}
	var fix_left int = 0
	var fix_right int = 0
	var float_left int = 0
	var float_right int = 0
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			if locked[i] == '1' {
				fix_left++
			} else {
				float_left++
			}
		} else {
			if locked[i] == '1' {
				fix_right++
			} else {
				float_right++
			}
		}
		if fix_right > fix_left+float_left+float_right {
			return false
		}
	}
	fix_left = 0
	fix_right = 0
	float_left = 0
	float_right = 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == '(' {
			if locked[i] == '1' {
				fix_left++
			} else {
				float_left++
			}
		} else {
			if locked[i] == '1' {
				fix_right++
			} else {
				float_right++
			}
		}
		if fix_left > fix_right+float_left+float_right {
			return false
		}
	}
	if fix_left > fix_right+float_left+float_right {
		return false
	}
	return true
}
