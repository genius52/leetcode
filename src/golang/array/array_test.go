package array

import "testing"

func Test_minOperations(t *testing.T) {
	logs := []string{"d1/","../","../","../"}
	res := minOperations(logs)
	if res != 0{
		t.Error("Test_minOperations fail")
	}
	t.Log("hello world")
}