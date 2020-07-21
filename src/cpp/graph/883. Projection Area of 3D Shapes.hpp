#include <vector>
#include <unordered_map>
using namespace std;

//Input: [[1,2],[3,4]]
//Output: 17
//Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j).
class Solution_883 {
public:
    int projectionArea(vector<vector<int>>& grid) {
        int cnt = grid.size();
        int res = 0;
        std::unordered_map<int,int> graph;
        for(int i = 0;i < cnt;i++){
            int max_x = 0;
            for(int j = 0;j < grid[i].size();j++){
                if(grid[i][j] > max_x)
                    max_x = grid[i][j];
                if(graph.find(j) == graph.end())
                    graph[j] = grid[i][j];
                else{
                    if(grid[i][j] > graph[j])
                        graph[j] = grid[i][j];
                }
                if(grid[i][j] > 0)
                    res++;
            }
            res += max_x;
        }
        for(auto it : graph){
            res += it.second;
        }
        return res;
    }
};