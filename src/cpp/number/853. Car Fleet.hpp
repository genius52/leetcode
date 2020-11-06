#include <vector>
#include <map>
using namespace std;

////Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
////Output: 3
////Explanation:
////The cars starting at 10 and 8 become a fleet, meeting each other at 12.
////The car starting at 0 doesn't catch up to any other car, so it is a fleet by itself.
////The cars starting at 5 and 3 become a fleet, meeting each other at 6.
////Note that no other cars meet these fleets before the destination, so the answer is 3.
class Solution_853 {
public:
    int carFleet(int target, vector<int>& position, vector<int>& speed) {
        int len = position.size();
        if (len <= 1)
            return len;
        std::map<int,float> record;//distance : duration
        for(int i = 0;i < len;i++){
            record[target - position[i]] = float((target - position[i]))/float(speed[i]);
        }
        int res = 1;
        auto it = record.begin();
        float pre_duration = it->second;
        it++;
        for(;it != record.end();it++){
            if(it->second <= pre_duration)
                continue;
            pre_duration = it->second;
            res++;
        }
        return res;
    }
};