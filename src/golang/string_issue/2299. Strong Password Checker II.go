package string_issue

//It has at least 8 characters.
//It contains at least one lowercase letter.
//It contains at least one uppercase letter.
//It contains at least one digit.
//It contains at least one special character. The special characters are the characters in the following string: "!@#$%^&*()-+".
//It does not contain 2 of the same character in adjacent positions (i.e., "aab" violates this condition, but "aba" does not).
func strongPasswordCheckerII(password string) bool {
	var l int = len(password)
	if l < 8{
		return false
	}
	var lower bool = false
	var upper bool = false
	var one_digit bool = false
	var special bool = false
	for i := 0;i < l;i++{
		if i > 0{
			if password[i] == password[i - 1]{
				return false
			}
		}
		if password[i] >= 'a' && password[i] <= 'z'{
			lower = true
		}
		if password[i] >= 'A' && password[i] <= 'Z'{
			upper = true
		}
		if password[i] >= '0' && password[i] <= '9'{
			one_digit = true
		}
		if password[i] == '!' || password[i] == '@' || password[i] == '#' || password[i] == '$' ||
			password[i] == '%' || password[i] == '^' || password[i] == '&' || password[i] == '*' ||
			password[i] == '(' || password[i] == ')' || password[i] == '-' || password[i] == '+'{
			special = true
		}
	}
	return lower && upper && special && one_digit
}