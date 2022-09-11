#include <vector>
#include <set>
using namespace std;

class Solution_2398 {
public:
    int maximumRobots(vector<int>& chargeTimes, vector<int>& runningCosts, long long budget) {
        int len = chargeTimes.size();
        std::multiset<int> s;
        int left = 0;
        int right = 0;
        int res = 0;
        long long cur_sum = 0;
        //max(chargeTimes) + k * sum(runningCosts)
        while (left < len && right < len){
            s.insert(chargeTimes[right]);
            cur_sum += runningCosts[right];
            if(*s.rbegin() + (right - left + 1) * cur_sum > budget){
                while(left <= right && *s.rbegin() + (right - left + 1) * cur_sum > budget){
                    s.erase(s.find(chargeTimes[left]));
                    cur_sum -= runningCosts[left];
                    left++;
                }
            }else{
                res = max(res,right - left + 1);
            }
            right++;
        }
        return res;
    }
};