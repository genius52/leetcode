package string_issue

//每个元音 'a' 后面都只能跟着 'e'
//每个元音 'e' 后面只能跟着 'a' 或者是 'i'
//每个元音 'i' 后面 不能 再跟着另一个 'i'
//每个元音 'o' 后面只能跟着 'i' 或者是 'u'
//每个元音 'u' 后面只能跟着 'a'
func countVowelPermutation(n int) int {
	pre_a := 1//end with 'a'
	pre_e := 1//end with 'e'
	pre_i := 1//end with 'i'
	pre_o := 1//end with 'o'
	pre_u := 1//end with 'u'
	MOD := int(1e9 + 7)
	for i := 1;i < n;i++{
		//a结尾
		cur_a := pre_e + pre_i + pre_u
		cur_a %= MOD
		//e结尾
		cur_e := pre_a + pre_i
		cur_e %= MOD
		//i结尾
		cur_i := pre_e + pre_o//i可以跟着e and o
		cur_i %= MOD
		//o结尾
		cur_o := pre_i
		cur_o %= MOD
		//u结尾
		cur_u := pre_i + pre_o
		cur_u %= MOD
		pre_a = cur_a
		pre_e = cur_e
		pre_i = cur_i
		pre_o = cur_o
		pre_u = cur_u
	}
	var res int = (pre_a + pre_e + pre_i + pre_o + pre_u) % MOD
	return res
}