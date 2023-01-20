#include <vector>
#include <queue>
using namespace std;

class Solution_2503 {
public:
    vector<int> maxPoints(vector<vector<int>>& grid, vector<int>& queries) {
        int rows = grid.size();
        int columns = grid[0].size();
        int l = queries.size();
        std::vector<std::pair<int,int>> record(l);
        for(int i = 0;i < l;i++){
            record[i] = {queries[i],i};
        }
        std::sort(record.begin(),record.end());
        std::vector<std::vector<bool>> visited(rows,std::vector<bool>(columns));
        std::vector<std::vector<int>> dirs{{-1,0},{1,0},{0,-1},{0,1}};
        std::vector<int> res(l);
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::vector<std::pair<int,std::pair<int,int>>>, std::greater<std::pair<int,std::pair<int,int>>>> to_be_visit;
        to_be_visit.push({grid[0][0],{0,0}});
        int pre_scores = 0;
        for(int i = 0;i < l;i++){
            int scores = 0;
            while (!to_be_visit.empty() && to_be_visit.top().first < record[i].first){
                auto cur = to_be_visit.top();
                to_be_visit.pop();
                visited[cur.second.first][cur.second.second] = true;
                scores++;
                for(auto dir : dirs){
                    auto x = cur.second.first + dir[0];
                    auto y = cur.second.second + dir[1];
                    if(x < 0 || x >= rows || y < 0 || y >= columns){
                        continue;
                    }
                    if(visited[x][y])
                        continue;
                    visited[x][y] = true;
                    to_be_visit.push({grid[x][y],{x,y}});
                }
            }
            res[record[i].second] = pre_scores + scores;
            pre_scores = res[record[i].second];
        }
        return res;
    }
};