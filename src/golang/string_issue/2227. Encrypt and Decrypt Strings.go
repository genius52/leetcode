package string_issue

type Encrypter struct {
	keys []byte
	char_idx map[uint8]int

	values []string
	string_idx map[string][]int

	allow map[string]int
}

//字符串 加密 按下述步骤进行：
//对字符串中的每个字符 c ，先从 keys 中找出满足 keys[i] == c 的下标 i 。
//在字符串中，用values[i] 替换字符 c 。

//字符串 解密 按下述步骤进行：
//将字符串每相邻 2 个字符划分为一个子字符串，对于每个子字符串 s ，找出满足 values[i] == s 的一个下标 i 。
//如果存在多个有效的 i ，从中选择 任意 一个。这意味着一个字符串解密可能得到多个解密字符串。
//在字符串中，用 keys[i] 替换 s 。

func Constructor2227(keys []byte, values []string, dictionary []string) Encrypter {
	var obj Encrypter
	return obj
}

func (this *Encrypter) Encrypt(word1 string) string {
	var res string
	for _,c := range word1{
		res += this.values[this.char_idx[uint8(c)]]
	}
	return res
}

func (this *Encrypter) Decrypt(word2 string) int {
	return 0
}