package number

//           1          | 0xxxxxxx
//            2          | 110xxxxx 10xxxxxx
//            3          | 1110xxxx 10xxxxxx 10xxxxxx
//            4          | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
func ValidUtf8(data []int) bool {
	var l int = len(data)
	for i := 0;i < l;{
		if (data[i] & 0b10000000) == 0{//0xxxxxxx
			i++
		}else if i < l - 1 && (data[i] & 0b11000000) == 0b11000000 && (data[i] | 0b00100000) != data[i] &&
			(data[i + 1] & 0b10000000) == 0b10000000 && (data[i + 1] | 0b01000000) != data[i + 1]{ //110xxxxx
			i += 2
		}else if i < l - 2 && (data[i] & 0b11100000) == 0b11100000 && (data[i] | 0b00010000) != data[i] &&
			(data[i + 1] & 0b10000000) == 0b10000000 && (data[i + 1] | 0b01000000) != data[i + 1] &&
			(data[i + 2] & 0b10000000) == 0b10000000 && (data[i + 2] | 0b01000000) != data[i + 2]{
			i += 3
		}else if i < l - 3 && (data[i] & 0b11110000) == 0b11110000 && (data[i] | 0b00001000) != data[i] &&
			(data[i + 1] & 0b10000000) == 0b10000000 && (data[i + 1] | 0b01000000) != data[i + 1] &&
			(data[i + 2] & 0b10000000) == 0b10000000 && (data[i + 2] | 0b01000000) != data[i + 2] &&
			(data[i + 3] & 0b10000000) == 0b10000000 && (data[i + 3] | 0b01000000) != data[i + 3] {
			i += 4
		}else{
			return false
		}
	}
	return true
}