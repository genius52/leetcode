#include <vector>
#include <set>
#include <unordered_map>
using namespace std;

//[1,0,2,0,3,0,2,0,0,0,1,2,3]
class Solution_1488 {
public:
    vector<int> avoidFlood(vector<int>& rains) {
        int len = rains.size();
        std::vector<int> res(len,1);
        std::set<int> zero_pos;//晴天的日期
        std::unordered_map<int,int> filled_lake;//每个湖上一次下雨的日期
        for (int i = 0; i < len; ++i) {
            if (rains[i] == 0){//当天不下雨，保存日期
                zero_pos.insert(i);
            }else{
                if (filled_lake.find(rains[i]) == filled_lake.end()){ //这个湖之前是空的，保存 id : 日期
                    filled_lake[rains[i]] = i;
                }else{
                    auto it = zero_pos.upper_bound(filled_lake[rains[i]]); //这个湖已经满了，找找上一次下雨之前有没有晴天
                    if (it == zero_pos.end()){
                        return std::vector<int>();
                    }else{
                        auto pre_zero = *it;
                        res[pre_zero] = rains[i];
                        zero_pos.erase(pre_zero);
                        filled_lake[rains[i]] = i;
                    }
                }
                res[i] = -1;
            }
        }
        return res;
    }
};