package number

//你的音乐播放器里有 N 首不同的歌，在旅途中，你的旅伴想要听 L 首歌（不一定不同，即，允许歌曲重复）。
//每首歌至少播放一次。
//一首歌只有在其他 K 首歌播放完之后才能再次播放。
//0 <= K < N <= L <= 100

//输入：N = 3, L = 3, K = 1
//输出：6
//解释：有 6 种可能的播放列表。[1, 2, 3]，[1, 3, 2]，[2, 1, 3]，[2, 3, 1]，[3, 1, 2]，[3, 2, 1].

//输入：N = 2, L = 3, K = 0
//输出：6
//解释：有 6 种可能的播放列表。[1, 1, 2]，[1, 2, 1]，[2, 1, 1]，[2, 2, 1]，[2, 1, 2]，[1, 2, 2]

//输入：N = 2, L = 3, K = 1
//输出：2
//解释：有 2 种可能的播放列表。[1, 2, 1]，[2, 1, 2]
//dp[i][j] denotes the solution of i songs with j different songs. So the final answer should be dp[L][N]
func dp_numMusicPlaylists(n int,l int,k int,total_song int,diff_song int,memo [101][101]int)int{
	if total_song == l{
		if diff_song < n{
			return 0
		}
		return 1
	}
	if memo[diff_song][total_song] != 0{
		return memo[diff_song][total_song]
	}
	var res int = 0
	if diff_song < n{
		//choose different song
		res += (n - diff_song) * dp_numMusicPlaylists(n,l,k,total_song + 1,diff_song + 1,memo)
		res = res % 1000000007
	}
	//choose exist song before,k should be considered
	if diff_song >= k{
		res += (diff_song - k) * dp_numMusicPlaylists(n,l,k,total_song + 1,diff_song,memo)
	}
	res = res % 1000000007
	memo[diff_song][total_song] = res
	return res
}

func NumMusicPlaylists(n int, l int, k int) int {
	//MOD := 1000000007
	//var dp [][]int = make([][]int,l)//dp[1][1]
	//for i := 0;i < l;i++{
	//	dp[i] = make([]int,n)
	//}
	var memo [101][101]int
	return dp_numMusicPlaylists(n,l,k,0,0,memo)
}