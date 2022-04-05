package string_issue

//字符串 加密 按下述步骤进行：
//对字符串中的每个字符 c ，先从 keys 中找出满足 keys[i] == c 的下标 i 。
//在字符串中，用values[i] 替换字符 c 。

//字符串 解密 按下述步骤进行：
//将字符串每相邻 2 个字符划分为一个子字符串，对于每个子字符串 s ，找出满足 values[i] == s 的一个下标 i 。
//如果存在多个有效的 i ，从中选择 任意 一个。这意味着一个字符串解密可能得到多个解密字符串。
//在字符串中，用 keys[i] 替换 s 。

type Encrypter struct {
	encode_map map[uint8]string
	allow_map map[string]int
}

func Constructor2227(keys []byte, values []string, dictionary []string) Encrypter {
	var obj Encrypter
	obj.encode_map = make(map[uint8]string)
	for idx,c := range keys{
		obj.encode_map[c] = values[idx]
	}
	obj.allow_map = make(map[string]int)
	for _,s := range dictionary{
		s2 := obj.Encrypt(s)
		if len(s2) > 0{
			obj.allow_map[s2]++
		}
	}
	return obj
}

func (this *Encrypter) Encrypt(word1 string) string {
	var res string
	for _,c := range word1{
		if val,ok := this.encode_map[byte(c)];ok{
			res += val
		}else{
			return ""
		}
	}
	return res
}

func (this *Encrypter) Decrypt(word2 string) int {
	if cnt,ok := this.allow_map[word2];ok{
		return cnt
	}else{
		return 0
	}
}