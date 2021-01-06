package diagram

//Input
//["CombinationIterator", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
//[["abc", 2], [], [], [], [], [], []]
//Output
//[null, "ab", true, "ac", true, "bc", false]
//
//Explanation
//CombinationIterator itr = new CombinationIterator("abc", 2);
//itr.next();    // return "ab"
//itr.hasNext(); // return True
//itr.next();    // return "ac"
//itr.hasNext(); // return True
//itr.next();    // return "bc"
//itr.hasNext(); // return False
type CombinationIterator struct {
	q []string
	pos int
}

func dfs_1286(characters string,cur_pos int,rest_len int,pre_str string,res *[]string){
	if rest_len < 0{
		return
	}
	if rest_len == 0{
		*res = append(*res,pre_str)
		return
	}
	if cur_pos >= len(characters){
		return
	}
	dfs_1286(characters,cur_pos + 1,rest_len - 1,pre_str + string(characters[cur_pos]),res)
	dfs_1286(characters,cur_pos + 1,rest_len,pre_str,res)
}

func Constructor1286(characters string, combinationLength int) CombinationIterator {
	var obj CombinationIterator
	dfs_1286(characters,0,combinationLength,"",&obj.q)
	return obj
}

func (this *CombinationIterator) Next() string {
	res := this.q[this.pos]
	this.pos++
	return res
}

func (this *CombinationIterator) HasNext() bool {
	return this.pos < len(this.q)
}