#include <vector>
#include <set>
#include <unordered_map>
using namespace std;

//[1,0,2,0,3,0,2,0,0,0,1,2,3]
class Solution_1488 {
public:
    vector<int> avoidFlood(vector<int>& rains) {
        int len = rains.size();
        std::vector<int> res;
        res.resize(len);
        std::set<int> zero_pos;
        std::unordered_map<int,int> filled_lake;
        for (int i = 0; i < len; ++i) {
            if (rains[i] == 0){
                zero_pos.insert(i);
            }else{
                if (filled_lake.find(rains[i]) == filled_lake.end()){
                    filled_lake[rains[i]] = i;
                }else{
                    auto it = zero_pos.upper_bound(filled_lake[rains[i]]);
                    if (it == zero_pos.end()){
                        return std::vector<int>();
                    }else{
                        auto pre_zero = *it;
                        res[pre_zero] = rains[i];
                        zero_pos.erase(pre_zero);
                        filled_lake[rains[i]] = i;
                    }
                    //filled_lake.erase(rains[i]);
                }
                res[i] = -1;
            }
        }
        for (int i = 0;i < len;i++){
            if (res[i] == 0){
                res[i] = 1;
            }
        }
        return res;
    }
};