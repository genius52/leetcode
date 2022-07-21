package tree

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

//Implement a trie with insert, search, and startsWith methods.
//
//Example:
//
//Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // returns true
//trie.search("app");     // returns false
//trie.startsWith("app"); // returns true
//trie.insert("app");
//trie.search("app");     // returns true


type Trie struct {
	record [26]*Trie
	is_word bool
}

/** Initialize your data structure here. */
func Constructor208() Trie {
	var obj Trie
	return obj
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	l := len(word)
	if l == 0{
		return
	}
	visit := this
	for i := 0;i < l;i++{
		if visit.record[word[i] - 'a'] == nil{
			var t Trie
			visit.record[word[i] - 'a'] = &t
		}
		visit = visit.record[word[i] - 'a']
	}
	visit.is_word = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	l := len(word)
	if l == 0{
		return this.is_word
	}
	var visit *Trie = this
	i := 0
	for i < l{
		if nil == visit.record[word[i] - 'a']{
			return false
		}
		visit = visit.record[word[i] - 'a']
		i++
	}
	return visit.is_word
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	l := len(prefix)
	if l == 0{
		return true
	}
	var visit *Trie = this
	i := 0
	for i < l{
		if nil == visit.record[prefix[i] - 'a']{
			return false
		}
		visit = visit.record[prefix[i] - 'a']
		i++
	}
	return true
}