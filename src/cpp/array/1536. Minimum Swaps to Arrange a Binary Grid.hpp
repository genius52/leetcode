#include <vector>
#include <map>
using namespace std;

class Solution_1536 {
public:
    int minSwaps(vector<vector<int>>& grid) {
        std::map<int,int> graph;//line number <-> zero count
        int rows = grid.size();
        int columns = grid[0].size();
        for(int i = 0;i < rows;i++){
            int l = 0;
            for(int j = columns - 1;j >= 0;j--){
                if(grid[i][j] == 0){
                    l++;
                }else{
                    break;
                }
            }
            graph[i] = l;
        }
        int res = 0;
        for (int i = 0; i < rows - 1; ++i) {
            int target = rows - 1 - i;
            int steps = 0;
            int j = i;
            bool find = false;
            for(;j < rows;j++){
                if(graph[j] >= target){
                    find = true;
                    break;
                }
            }
            if(!find)
                return -1;
            res += j - i;
            for(int k = j;k > i;k--){
                graph[k] = graph[k - 1];
            }
        }
        return res;
    }
};