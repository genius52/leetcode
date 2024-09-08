package number

func generateKey(num1 int, num2 int, num3 int) int {
	var data [4]int
	var offset int = 0
	for offset < 4 {
		m1 := num1 % 10
		m2 := num2 % 10
		m3 := num3 % 10
		data[offset] = min(m1, m2, m3)
		num1 /= 10
		num2 /= 10
		num3 /= 10
		offset++
	}
	return data[3]*1000 + data[2]*100 + data[1]*10 + data[0]
}
