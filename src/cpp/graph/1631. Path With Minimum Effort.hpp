#include <vector>
#include <queue>
using namespace std;

class Solution_1631 {
public:
    int minimumEffortPath(vector<vector<int>>& heights) {
        int rows = heights.size();
        int columns = heights[0].size();
        std::vector<std::vector<int>> weights(rows,std::vector<int>(columns,2147483647));
        weights[0][0] = 0;
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::vector<std::pair<int,std::pair<int,int>>>, std::greater<std::pair<int,std::pair<int,int>>>> q;//weight,{row,column}
        q.push({0,{0,0}});
        std::vector<std::vector<int>> dirs{{-1,0},{1,0},{0,-1},{0,1}};
        while (!q.empty()){
            auto cur = q.top();
            q.pop();
            if(cur.second.first == rows - 1 && cur.second.second == columns - 1)
                break;
            for (auto dir : dirs) {
                std::pair<int,std::pair<int,int>> next;
                next.second.first = cur.second.first + dir[0];
                next.second.second = cur.second.second + dir[1];
                if(next.second.first < 0 || next.second.first >= rows || next.second.second < 0 || next.second.second >= columns)
                    continue;
                next.first = max(cur.first,abs(heights[next.second.first][next.second.second] - heights[cur.second.first][cur.second.second]));
                if (next.first >= weights[next.second.first][next.second.second])
                    continue;
                weights[next.second.first][next.second.second] = next.first;
                q.push(next);
            }
        }
        return weights[rows - 1][columns - 1];
    }
};