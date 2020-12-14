package diagram

import "math"

//Input: books = [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]], shelf_width = 4
//Output: 6
func dp_minHeightShelves(books [][]int,pos int,curlevel_width int,curlevel_height int,shelf_width int,memo *[1001][1001]int)int{
	var l int = len(books)
	if pos >= l{
		return curlevel_height
	}
	if memo[pos][curlevel_width] != 0{
		return memo[pos][curlevel_width]
	}
	if curlevel_width + books[pos][0] <= shelf_width{
		var total_jump_next int = math.MaxInt32
		if curlevel_width != 0{
			next_level_height := dp_minHeightShelves(books,pos,0,0,shelf_width,memo)
			total_jump_next = curlevel_height + next_level_height
		}
		if curlevel_height < books[pos][1]{
			curlevel_height = books[pos][1]
		}
		total_not_jump := dp_minHeightShelves(books,pos + 1,curlevel_width + books[pos][0],curlevel_height,shelf_width,memo)
		if total_not_jump < total_jump_next{
			memo[pos][curlevel_width] = total_not_jump
		}else{
			memo[pos][curlevel_width] = total_jump_next
		}
		return memo[pos][curlevel_width]
	}else{
		next_level_height := dp_minHeightShelves(books,pos,0,0,shelf_width,memo)
		memo[pos][curlevel_width] = curlevel_height + next_level_height
		return memo[pos][curlevel_width]
	}
}

func MinHeightShelves(books [][]int, shelf_width int) int {
	var memo [1001][1001]int
	res := dp_minHeightShelves(books,0,0,0,shelf_width,&memo)
	return res
}