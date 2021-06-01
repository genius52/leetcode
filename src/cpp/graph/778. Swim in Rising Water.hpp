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

//    int dfs_swimInWater(vector<vector<int>>& grid,int rows,int columns,int r,int c,int last_height,int cur_time,
//                        std::vector<std::vector<int>>& visited,std::vector<std::vector<std::vector<int>>>& memo){
//        if(r < 0 || r >= rows || c < 0 || c >= columns){
//            return -1;
//        }
//        if(r == rows - 1 && c == columns - 1){
//            if(last_height >= grid[r][c])
//                return 0;
//            else{
//                return grid[r][c] - last_height;
//            }
//        }
//        if(visited[r][c])
//            return -1;
//        if(memo[r][c][cur_time] != 0)
//            return memo[r][c][cur_time];
//        visited[r][c] = true;
//        int res = 0;
//        int wait_time = 0;
//        if(r != 0 || c != 0){
//            if(last_height != grid[r][c]) {
//                int high = max(grid[r][c],last_height);
//                if(high > cur_time){
//                    wait_time = int(abs(high - cur_time));
//                }
//                cur_time += wait_time;
//            }
//        }
//
//        int next_steps = 2147483647;
//        int steps1 = dfs_swimInWater(grid,rows,columns,r + 1,c,max(grid[r][c],last_height),cur_time,visited,memo);
//        if(steps1 != -1){
//            next_steps = min(next_steps,steps1);
//        }
//        int steps2 = dfs_swimInWater(grid,rows,columns,r - 1,c,max(grid[r][c],last_height),cur_time,visited,memo);
//        if(steps2 != -1){
//            next_steps = min(next_steps,steps2);
//        }
//        int steps3 = dfs_swimInWater(grid,rows,columns,r,c - 1,max(grid[r][c],last_height),cur_time,visited,memo);
//        if(steps3 != -1){
//            next_steps = min(next_steps,steps3);
//        }
//        int steps4 = dfs_swimInWater(grid,rows,columns,r,c + 1,max(grid[r][c],last_height),cur_time,visited,memo);
//        if(steps4 != -1){
//            next_steps = min(next_steps,steps4);
//        }
//        if(next_steps == 2147483647){
//            res = -1;
//        }else{
//            res = next_steps + wait_time;
//        }
//        visited[r][c] = false;
//        memo[r][c][cur_time] = res;
//        return res;
//    }
//
//    int swimInWater(vector<vector<int>>& grid){
//        int rows = grid.size();
//        int columns = grid[0].size();
//        std::vector<std::vector<int>> visited(rows,std::vector<int>(columns));
//        std::vector<std::vector<std::vector<int>>> memo(rows,std::vector<std::vector<int>>(columns,std::vector<int>(rows * columns)));
//        return dfs_swimInWater(grid,rows,columns,0,0,-1,0,visited,memo);
//    }
};