#include <vector>
#include <queue>
using namespace std;

class Solution_2662 {
public:
    int minimumCost(vector<int>& start, vector<int>& target, vector<vector<int>>& specialRoads) {
        int len = specialRoads.size();
        std::vector<int> from_start(len,1 << 31 - 1);
        std::priority_queue<std::pair<int,int>> small_top;
        for(int i = 0;i < len;i++){
            from_start[i] = min(abs(start[0] - specialRoads[i][2]) + abs(start[1] - specialRoads[i][3]),
            abs(start[0] - specialRoads[i][0]) + abs(start[1] - specialRoads[i][1]) + specialRoads[i][4]);
            small_top.push({from_start[i],i});
        }
        int res = abs(start[0] - target[0]) + abs(start[1] - target[1]);
        while (!small_top.empty()){
            auto top = small_top.top();
            int cost_fromstart = top.first;
            int idx = top.second;
            small_top.pop();
            res = min(res,cost_fromstart + abs(target[0] - specialRoads[idx][2]) + abs(target[1] - specialRoads[idx][3]));
            for(int i = 0;i < len;i++){
//                from_start[i] = min(from_start[i], cost_fromstart + abs(specialRoads[idx][0] - specialRoads[i][0]) +
//                        abs(specialRoads[idx][1] - specialRoads[i][1]));
                int next_cost = cost_fromstart +  abs(specialRoads[idx][2] - specialRoads[i][0]) +
                        abs(specialRoads[idx][3] - specialRoads[i][1]) + specialRoads[i][4];
                if(next_cost < from_start[i]){
                    from_start[i] = next_cost;
                    small_top.push({from_start[i],i});
                }
            }
        }
        return res;
    }
};