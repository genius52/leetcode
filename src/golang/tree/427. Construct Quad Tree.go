package tree

type Node3 struct {
	Val bool
	IsLeaf bool
	TopLeft *Node3
	TopRight *Node3
	BottomLeft *Node3
    BottomRight *Node3
}


func recursive_construct(grid [][]int,r1 int,c1 int,r2 int,c2 int)*Node3{
	if r1 > r2 || c1 > c2{
		return nil
	}
	var is_leaf bool = true
	var cur_val int = grid[r1][c1]
loop:
	for i := r1;i <= r2;i++{
		for j := c1;j <= c2;j++{
			if grid[i][j] != cur_val{
				is_leaf = false
				break loop
			}
		}
	}
	if is_leaf{
		return &Node3{grid[r1][c1] == 1,true,nil,nil,nil,nil}
	}
	mid_r := (r1 + r2)/2
	mid_c := (c1 + c2)/2
	var cur Node3
	cur.TopLeft = recursive_construct(grid,r1,c1,mid_r,mid_c)
	cur.TopRight = recursive_construct(grid,r1,mid_c + 1,mid_r,c2)
	cur.BottomLeft = recursive_construct(grid,mid_r + 1,c1,r2,mid_c)
	cur.BottomRight = recursive_construct(grid,mid_r + 1,mid_c + 1,r2,c2)

	//if cur.TopLeft.Val == cur.TopRight.Val && cur.TopLeft.Val == cur.BottomLeft.Val && cur.TopLeft.Val == cur.BottomRight.Val{
	//	cur.Val = cur.TopLeft.Val
	//}
	return &cur
}

func Construct(grid [][]int) *Node3 {
	var l int = len(grid)
	return recursive_construct(grid,0,0,l - 1,l - 1)
}