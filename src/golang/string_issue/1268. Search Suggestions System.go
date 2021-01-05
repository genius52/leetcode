package string_issue

//Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
//Output: [
//["mobile","moneypot","monitor"],
//["mobile","moneypot","monitor"],
//["mouse","mousepad"],
//["mouse","mousepad"],
//["mouse","mousepad"]
//]
type dict_tree struct{
	chars [26]*dict_tree
	is_word bool
}

func dfs_suggestedProducts(pre string,node *dict_tree,cur *[]string){
	if node == nil{
		return
	}
	if len(*cur) == 3{
		return
	}
	if node.is_word{
		*cur = append(*cur,pre)
	}
	for i := 0;i < 26;i++{
		if node.chars[i] == nil{
			continue
		}
		var s string = pre + string(i + 'a')
		dfs_suggestedProducts(s,node.chars[i],cur)
	}
}

func SuggestedProducts(products []string, searchWord string) [][]string {
	var res [][]string
	var root dict_tree
	var word_cnt int = len(products)
	for i := 0;i < word_cnt;i++{
		var visit int = 0
		var word_len int = len(products[i])
		var node *dict_tree = &root
		for visit < word_len{
			find := node.chars[products[i][visit] - 'a']
			if find == nil{
				var w dict_tree
				node.chars[products[i][visit] - 'a'] = &w
				node = &w
			}else{
				node = node.chars[products[i][visit] - 'a']
			}
			if visit == word_len - 1{
				node.is_word = true
			}
			visit++
		}
	}
	var l int = len(searchWord)
	var node *dict_tree = &root
	var pre string
	var i int = 0
	for ;i < l;i++{
		find := node.chars[searchWord[i] - 'a']
		var cur []string
		if find != nil{
			pre += string(searchWord[i])
			node = node.chars[searchWord[i] - 'a']

			dfs_suggestedProducts(pre,node,&cur)
		}else{
			break
		}
		res = append(res,cur)
	}
	for i < l{
		res = append(res,[]string{})
		i++
	}
	return res
}