#include <vector>
#include <queue>
using namespace std;

//1 下一步往右走，也就是你会从 grid[i][j]走到 grid[i][j + 1]
//2 下一步往左走，也就是你会从 grid[i][j]走到 grid[i][j - 1]
//3 下一步往下走，也就是你会从 grid[i][j]走到 grid[i + 1][j]
//4 下一步往上走，也就是你会从 grid[i][j]走到 grid[i - 1][j]

class Solution_1368 {
public:
    //dijkstra solution
    int minCost(vector<vector<int>>& grid) {
        std::vector<std::vector<int>> dirs{{0,1},{0,-1},{1,0},{-1,0}};
        int rows = grid.size();
        int columns = grid[0].size();
        std::vector<std::vector<int>> graph(rows,std::vector<int>(columns,2147483647));
        graph[0][0] = 0;
        //std::priority_queue<std::pair<int,std::pair<int,int>>> q;//{cost,{row,column}}
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::vector<std::pair<int,std::pair<int,int>>>, std::greater<std::pair<int,std::pair<int,int>>>> q;
        q.push({0,{0,0}});
        std::unordered_set<int> visited;
        while (!q.empty()){
            auto cur = q.top();
            q.pop();
            auto cur_cost = cur.first;
            auto r = cur.second.first;
            auto c = cur.second.second;
            if(graph[r][c] < cur_cost){
                continue;
            }
            auto cost_zero_r = r + dirs[grid[r][c] - 1][0];
            auto cost_zero_c = c + dirs[grid[r][c] - 1][1];
            for(int i = 0;i < 4;i++){
                auto next_r = r + dirs[i][0];
                auto next_c = c + dirs[i][1];
                if(next_r >= 0 && next_r < rows && next_c >= 0 && next_c < columns){
                    if(next_r == cost_zero_r && next_c == cost_zero_c){
                        if(cur_cost < graph[next_r][next_c]){
                            graph[next_r][next_c] = cur_cost;
                            q.push({cur_cost,{next_r,next_c}});
                        }
                    }else{
                        if((cur_cost + 1) < graph[next_r][next_c]){
                            graph[next_r][next_c] = cur_cost + 1;
                            q.push({cur_cost + 1,{next_r,next_c}});
                        }
                    }
                }
            }
        }
        return graph[rows - 1][columns - 1];
    }
};