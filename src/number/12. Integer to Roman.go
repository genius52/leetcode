package number

//12
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//Input: 58
//Output: "LVIII"
//Explanation: L = 50, V = 5, III = 3.
//Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
func cal_intToRoman(num int,record map[int]string) string{
	if num == 0{
		return ""
	}
	var cnt int = 0
	if num >= 1000{
		cnt = num - num % 1000
	}else if num >= 100 && num < 1000{
		cnt = num - num % 100
	}else if num >= 10 && num < 100{
		cnt = num - num % 10
	}else{
		cnt = num
	}
	var res string = record[cnt]
	res += cal_intToRoman(num - cnt,record)
	return res
}

func IntToRoman(num int) string {
	var record map[int]string = make(map[int]string)
	record[1] = "I"
	record[2] = "II"
	record[3] = "III"
	record[4] = "IV"
	record[5] = "V"
	record[6] = "VI"
	record[7] = "VII"
	record[8] = "VIII"
	record[9] = "IX"
	record[10] = "X"
	record[20] = "XX"
	record[30] = "XXX"
	record[40] = "XL"
	record[50] = "L"
	record[60] = "LX"
	record[70] = "LXX"
	record[80] = "LXXX"
	record[90] = "XC"
	record[100] = "C"
	record[200] = "CC"
	record[300] = "CCC"
	record[400] = "CD"
	record[500] = "D"
	record[600] = "DC"
	record[700] = "DCC"
	record[800] = "DCCC"
	record[900] = "CM"
	record[1000] = "M"
	record[2000] = "MM"
	record[3000] = "MMM"
	res := cal_intToRoman(num,record)
	return res
}