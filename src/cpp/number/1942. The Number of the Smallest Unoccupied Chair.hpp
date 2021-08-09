#include <vector>
#include <set>
#include <map>
#include <queue>
using namespace std;

class Solution_1942 {
public:
    int smallestChair(vector<vector<int>>& times, int targetFriend) {
        int len = times.size();
        int target_arrival = times[targetFriend][0];
        std::sort(times.begin(),times.end(),[](const std::vector<int>& v1, const vector<int>& v2)->bool
            {
                return v1[0] < v2[0];
            });
        std::priority_queue<std::pair<int,int>,vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        std::set<int> can_use;
        for(int i = 0;i < len;i++){
            can_use.insert(i);
        }
        int res = 0;
        for(int i = 0;i < len;i++){
            while(q.size() > 0){//release the seat before current arrive time
                auto top = q.top();
                if(top.first <= times[i][0]){
                    can_use.insert(top.second);
                    q.pop();
                }else{
                    break;
                }
            }
            //choose the minimum seat
            res = *can_use.begin();
            q.push({times[i][1],res});
            can_use.erase(can_use.begin());
            if(times[i][0] == target_arrival){
                break;
            }
        };
        return res;
    }
};