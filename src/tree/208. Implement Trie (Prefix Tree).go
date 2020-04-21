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
	tries := this
	for i := 0;i < l;i++{
		if tries.record[word[i] - 'a'] == nil{
			var t Trie
			if i == l - 1{
				t.is_word = true
			}
			tries.record[word[i] - 'a'] = &t
		}else{
			if i == l - 1{
				tries.record[word[i] - 'a'].is_word = true
			}
		}
		tries = tries.record[word[i] - 'a']
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	l := len(word)
	if l == 0{
		if this.is_word{
			return true
		}
		return false
	}
	var word_dict *Trie = this
	i := 0
	for ;i < l;i++ {
		if nil == word_dict.record[word[i] - 'a']{
			return false
		}
		word_dict = word_dict.record[word[i] - 'a']
	}
	if i != l{
		return false
	}
	return word_dict.is_word
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	l := len(prefix)
	if l == 0{
		return true
	}
	var word_dict *Trie = this
	i := 0
	for ;i < l;i++ {
		if nil == word_dict.record[prefix[i] - 'a']{
			return false
		}
		word_dict = word_dict.record[prefix[i] - 'a']
	}
	return true
}