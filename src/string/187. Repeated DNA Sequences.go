package string

//All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG". When studying DNA,
//it is sometimes useful to identify repeated sequences within the DNA.
//Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.
//Example:
//Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
//Output: ["AAAAACCCCC", "CCCCCAAAAA"]
func findRepeatedDnaSequences(s string) []string {
	var res []string
	l := len(s)
	if l < 10{
		return res
	}
	var record map[string]int = make(map[string]int)
	for i := 0;i < l - 9;i++{
		sub := s[i:i+10]
		if _,ok := record[sub];ok{
			record[sub]++
		}else{
			record[sub] = 1
		}
	}
	for k,v := range record{
		if v > 1{
			res = append(res,k)
		}
	}
	return res
}
