#include <vector>
#include <queue>
using namespace std;

class Solution_2290 {
public:
//    void minimumObstacles(vector<vector<int>>& grid,int rows,int columns,int r,int c,std::vector<std::vector<int>> dp){
//        if(r < 0 || r >= rows || c < 0 || c >= columns || grid[r][c] == 1 || dp[r][c] == 0)
//            return;
//        dp[r][c] = 0;
//        minimumObstacles(grid,rows,columns,r - 1,c,dp);
//        minimumObstacles(grid,rows,columns,r + 1,c,dp);
//        minimumObstacles(grid,rows,columns,r,c - 1,dp);
//        minimumObstacles(grid,rows,columns,r,c + 1,dp);
//    }

    int minimumObstacles(vector<vector<int>>& grid) {
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::deque<std::pair<int,std::pair<int,int>>>, std::greater<std::pair<int,std::pair<int,int>>>> q;//count,{x,y}
        int rows = grid.size();
        int columns = grid[0].size();
        std::vector<std::vector<int>> dp(rows,std::vector<int>(columns,-1));
        std::vector<std::vector<bool>> visited(rows,std::vector<bool>(columns));
        dp[0][0] = 0;
        visited[0][0] = true;
        //minimumObstacles(grid,rows,columns,0,0,dp);
        q.push({0,{0,0}});
        std::vector<std::pair<int,int>> dirs{{-1,0},{1,0},{0,-1},{0,1}};
        while(!q.empty()){
            auto cur = q.top();
            q.pop();
            for(auto dir : dirs){
                int next_r = cur.second.first + dir.first;
                int next_c = cur.second.second + dir.second;
                int next_cost = cur.first;
                if(next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns || visited[next_r][next_c])
                    continue;
                if(grid[next_r][next_c] == 1)
                    next_cost++;
//                if(next_cost >= dp[next_r][next_c])
//                    continue;
                dp[next_r][next_c] = next_cost;
                visited[next_r][next_c] = true;
                q.push({next_cost,{next_r,next_c}});
            }
        }
        return dp[rows - 1][columns - 1];
    }
};