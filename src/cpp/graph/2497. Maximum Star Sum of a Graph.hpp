#include <vector>
#include <queue>
using namespace std;

class Solution_2497 {
public:
    int maxStarSum(vector<int>& vals, vector<vector<int>>& edges, int k) {
        int len = vals.size();
        std::vector<std::priority_queue<int, std::deque<int>, std::greater<int> >> record(len);
        for(auto edge : edges){
            if(vals[edge[1]] > 0){
                if(record[edge[0]].size() < k)
                    record[edge[0]].push(vals[edge[1]]);
                else{
                    if(vals[edge[1]] > record[edge[0]].top()){
                        record[edge[0]].push(vals[edge[1]]);
                        record[edge[0]].pop();
                    }
                }
            }
            if(vals[edge[0]] > 0){
                if(record[edge[1]].size() < k)
                    record[edge[1]].push(vals[edge[0]]);
                else{
                    if(vals[edge[0]] > record[edge[1]].top()){
                        record[edge[1]].push(vals[edge[0]]);
                        record[edge[1]].pop();
                    }
                }
            }
        }
        int res = -2147483648;
        for(int i = 0;i < len;i++){
            int cur = vals[i];
            while (!record[i].empty()){
                cur += record[i].top();
                record[i].pop();
            }
            if(cur > res)
                res = cur;
        }
        return res;
    }
};