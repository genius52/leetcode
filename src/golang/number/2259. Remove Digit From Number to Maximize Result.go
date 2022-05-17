package number

import "sort"

func removeDigit(number string, digit byte) string {
	var l int = len(number)
	var record []string
	for i := 0;i < l;i++{
		if number[i] == digit{
			record = append(record,number[:i] + number[i + 1:])
		}
	}
	if len(record) == 0{
		return number
	}
	sort.Slice(record, func(i, j int) bool {
		return record[i] > record[j]
	})
	return record[0]
}

func removeDigit2(number string, digit byte) string{
	var l int = len(number)
	var idx int = -1
	for i := 0;i < l;i++{
		if number[i] == digit{
			if i != l - 1 && number[i] < number[i + 1]{
				return number[:i] + number[i + 1:]
			}
			idx = i
		}
	}
	if idx == -1{
		return number
	}
	return number[:idx] + number[idx + 1:]
}