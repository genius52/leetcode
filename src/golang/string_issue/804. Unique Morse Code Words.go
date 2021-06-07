package string_issue

//[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
func uniqueMorseRepresentations(words []string) int {
	var morse [26]string = [26]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	var record map[string]bool = make(map[string]bool)
	for _,word := range words{
		var s string
		for _,c := range word{
			s += morse[c - 'a']
		}
		record[s] = true
	}
	return len(record)
}