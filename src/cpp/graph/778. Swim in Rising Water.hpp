#include <vector>
#include <queue>
#include <array>
using namespace std;

//Input: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
//Output: 16
//Explanation:
// 0  1  2  3  4
//24 23 22 21  5
//12 13 14 15 16
//11 17 18 19 20
//10  9  8  7  6
//
//The final route is marked in bold.
//We need to wait until time 16 so that (0, 0) and (4, 4) are connected.
class Solution_778 {
public:
    int swimInWater(vector<vector<int>>& grid){
        int rows = grid.size();
        int columns = grid[0].size();
        std::vector<std::vector<bool>> visited(rows,std::vector<bool>(columns));
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::deque<std::pair<int,std::pair<int,int>>>, std::greater<std::pair<int,std::pair<int,int>>>> q;
        q.push({grid[0][0],{0,0}});
        std::array<std::array<int,2>,4> dirs{{{-1,0},{1,0},{0,-1},{0,1}}};
        int res = 0;
        while(q.size() > 0){
            auto top = q.top();
            q.pop();
            visited[top.second.first][top.second.second] = true;
            res = max(res,top.first);
            if(top.second.first == (rows - 1) && top.second.second == (columns - 1)){
                return res;
            }
            for(auto d : dirs){
                int i = top.second.first + d[0];
                int j = top.second.second + d[1];
                if(i < 0 || i >= rows || j < 0 || j >= columns)
                    continue;
                if(visited[i][j])
                    continue;
                std::pair<int,std::pair<int,int>> p;
                p.first = max(p.first,grid[i][j]);
                p.second.first = i;
                p.second.second = j;
                q.push(p);
            }
        }
        return 1;
    }

    std::vector<std::vector<int>> dirs{{-1,0},{1,0},{0,-1},{0,1}};
    void dfs_swimInWater(vector<vector<int>>& grid,int rows,int columns,int r,int c,int last_height,std::vector<std::vector<int>>& dp){
        if(r < 0 || r >= rows || c < 0 || c >= columns){
            return;
        }
        if(max(grid[r][c],last_height) >= dp[r][c])
            return;

        dp[r][c] = max(last_height,grid[r][c]);

        for(auto d : dirs){
            int i = r + d[0];
            int j = c + d[1];
            dfs_swimInWater(grid,rows,columns,i,j,dp[r][c],dp);
        }
    }

    int swimInWater2(vector<vector<int>>& grid){
        int rows = grid.size();
        int columns = grid[0].size();
        std::vector<std::vector<int>> dp(rows,std::vector<int>(columns,2147483647));
        dfs_swimInWater(grid,rows,columns,0,0,grid[0][0],dp);
        return dp[rows - 1][columns - 1];
    }
};