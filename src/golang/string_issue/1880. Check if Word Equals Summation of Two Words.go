package string_issue

func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var l1 int = len(firstWord)
	var l2 int = len(secondWord)
	var l3 int = len(targetWord)
	var v1 int = 0
	for i := 0;i < l1;i++{
		v1 *= 10
		v1 += int(firstWord[i] - 'a')
	}
	var v2 int = 0
	for i := 0;i < l2;i++{
		v2 *= 10
		v2 += int(secondWord[i] - 'a')
	}
	var v3 int = 0
	for i := 0;i < l3;i++{
		v3 *= 10
		v3 += int(targetWord[i] - 'a')
	}
	return (v1 + v2) == v3
}