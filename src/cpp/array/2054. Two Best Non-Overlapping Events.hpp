#include <vector>
#include <queue>
using namespace std;

class Solution_2054 {
public:
    int maxTwoEvents(vector<vector<int>>& events) {
        std::sort(events.begin(),events.end(),[](const std::vector<int>& v1,const std::vector<int>& v2){
           return v1[0] < v2[0];
        });
        int len = events.size();
        int res = 0;
        int pre_max = 0;
        std::priority_queue<std::pair<int,int>,std::deque<std::pair<int,int>>,
        std::greater<std::pair<int,int>>> q;//small top,end_time,val
        for (int i = 0; i < len; ++i) {
            if (q.empty()){
                q.push({events[i][1],events[i][2]});
            }else{
                while (!q.empty()){
                    auto top = q.top();
                    if (top.first >= events[i][0]){
                        break;
                    }else{
                        pre_max = max(pre_max,top.second);
                        q.pop();
                    }
                }
                q.push({events[i][1],events[i][2]});
            }
            res = max(res,pre_max + events[i][2]);
        }
        return res;
    }
};