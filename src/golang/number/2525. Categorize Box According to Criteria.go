package number

func categorizeBox(length int, width int, height int, mass int) string {
	var volumn int64 = int64(length * width * height)
	var kinds []string
	if length >= 1e4 || width >= 1e4 || height >= 1e4 || volumn >= 1e9{
		kinds = append(kinds,"Bulky")
	}
	if mass >= 100{
		kinds = append(kinds,"Heavy")
	}
	if len(kinds) == 2{
		return "Both"
	}
	if len(kinds) == 0{
		return "Neither"
	}
	return kinds[0]
}