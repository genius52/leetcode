package number

func DoesValidArrayExist(derived []int) bool {
	var l int = len(derived)
	var data []int = make([]int, l)
	data[0] = 0
	for i := 0; i < l-1; i++ {
		data[i+1] = data[i] ^ derived[i]
	}
	if (data[0] ^ data[l-1]) == derived[l-1] {
		return true
	}
	data[0] = 1
	for i := 0; i < l-1; i++ {
		data[i+1] = data[i] ^ derived[i]
	}
	return (data[0] ^ data[l-1]) == derived[l-1]
}
