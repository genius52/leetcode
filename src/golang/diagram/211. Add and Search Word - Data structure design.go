package diagram

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

type WordDictionary struct {
	record [26]*WordDictionary
	is_word bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	var obj WordDictionary
	return obj
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	l := len(word)
	if l == 0{
		return
	}
	var word_dict *WordDictionary = this
	for i := 0;i < l;i++{
		if nil == word_dict.record[word[i] - 'a']{
			var d WordDictionary
			if i == l - 1{
				d.is_word = true
			}
			word_dict.record[word[i] - 'a'] = &d
		}
		word_dict = word_dict.record[word[i] - 'a']
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	l := len(word)
	if l == 0{
		if this.is_word{
			return true
		}
		return false
	}
	var word_dict *WordDictionary = this
	i := 0
	for ;i < l;i++ {
		if word[i] == '.'{
			var find bool = false
			for _,d := range word_dict.record{
				if nil == d{
					continue
				}
				find = d.Search(word[i + 1 :])
				if find{
					return true
				}
			}
			if !find{
				return false
			}
		}else{
			if nil == word_dict.record[word[i] - 'a']{
				return false
			}
			word_dict = word_dict.record[word[i] - 'a']
		}
	}
	if i != l{
		return false
	}
	return word_dict.is_word
}