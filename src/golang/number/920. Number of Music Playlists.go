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
func dp_numMusicPlaylists(n int,target int,k int,total_choosed int,diff_choosed int,memo *[101][101]int)int{
	if total_choosed == target{
		if diff_choosed < n{
			return 0
		}
		return 1
	}
	if target - total_choosed < n - diff_choosed{
		return 0
	}
	if memo[total_choosed][diff_choosed] != -1{
		return memo[total_choosed][diff_choosed]
	}
	var res int = 0
	var MOD int = 1e9 + 7
	//choose different song
	res += (n - diff_choosed) * dp_numMusicPlaylists(n,target,k,total_choosed + 1,diff_choosed + 1,memo) % MOD
	//choose exist song before,k should be considered
	res += max_int(0,diff_choosed - k) * dp_numMusicPlaylists(n,target,k,total_choosed + 1,diff_choosed,memo) % MOD
	memo[total_choosed][diff_choosed] = res
	return res
}

func NumMusicPlaylists(n int, l int, k int) int {
	var memo [101][101]int
	for i := 0;i < 101;i++{
		for j := 0;j < 101;j++{
			memo[i][j] = -1
		}
	}
	return dp_numMusicPlaylists(n,l,k,0,0,&memo)
}