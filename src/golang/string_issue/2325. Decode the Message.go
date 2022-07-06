package string_issue

func DecodeMessage(key string, message string) string {
	var visited [26]bool
	var m map[uint8]uint8 = make(map[uint8]uint8)
	var idx uint8 = 0
	for i := 0;i < len(key);i++{
		if key[i] == ' ' || visited[key[i] - 'a']{
			continue
		}
		visited[key[i] - 'a'] = true
		m[key[i]] = 'a' + idx
		idx++
	}
	var res string
	for i := 0;i < len(message);i++{
		if message[i] == ' '{
			res += " "
		}else{
			res += string(m[message[i]])
		}
	}
	return res
}