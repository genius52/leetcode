#include <vector>
#include <queue>
using namespace std;

class Solution_2812 {
public:
    int maximumSafenessFactor(vector<vector<int>>& grid) {
        int rows = grid.size();
        int columns = grid[0].size();
        if(grid[0][0] == 1 || grid[rows - 1][columns - 1] == 1)
            return 0;
        std::queue<std::pair<int,int>> q1;
        for(int i = 0;i < rows;i++){
            for(int j = 0;j < columns;j++){
                if(grid[i][j] == 1){
                    q1.push({i,j});
                }
            }
        }
        std::vector<std::pair<int,int>> dirs{{-1,0},{1,0},{0,-1},{0,1}};
        int distance = 2;
        while (!q1.empty()){
            int cur_len = q1.size();
            for(int i = 0;i < cur_len;i++){
                auto cur = q1.front();
                q1.pop();
                for(auto dir : dirs){
                    std::pair<int,int> next;
                    next.first = cur.first + dir.first;
                    next.second = cur.second + dir.second;
                    if(next.first < 0 || next.first >= rows || next.second < 0 || next.second >= columns){
                        continue;
                    }
                    if(grid[next.first][next.second] != 0 && grid[next.first][next.second] <= distance)
                        continue;
                    grid[next.first][next.second] = distance;
                    q1.push({next.first,next.second});
                }
            }
            distance++;
        }
        std::vector<std::vector<bool>> visited(rows,std::vector<bool>(columns));
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::vector<std::pair<int,std::pair<int,int>>>, std::less<std::pair<int,std::pair<int,int>>>> q2;
        q2.push({grid[0][0],{0,0}});
        visited[0][0] = true;
        int min_distance = grid[0][0];
        bool find = false;
        while (!q2.empty()){
            auto cur = q2.top();
            if(grid[cur.second.first][cur.second.second] < min_distance)
                min_distance = grid[cur.second.first][cur.second.second];
            if(cur.second.first == rows - 1 && cur.second.second == columns - 1){
                find = true;
                break;
            }
            q2.pop();
            for(auto dir : dirs){
                std::pair<int,std::pair<int,int>> next;
                next.second.first = cur.second.first + dir.first;
                next.second.second = cur.second.second + dir.second;
                if(next.second.first < 0 || next.second.first >= rows || next.second.second < 0 || next.second.second >= columns ||
                   visited[next.second.first][next.second.second] || grid[next.second.first][next.second.second] == 1){
                    continue;
                }

                visited[next.second.first][next.second.second] = true;
                q2.push({grid[next.second.first][next.second.second],{next.second.first,next.second.second}});
            }
        }
        if(!find)
            return 0;
        return min_distance - 1;
    }
};