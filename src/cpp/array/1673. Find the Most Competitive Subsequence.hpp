#include <vector>
//#include <queue>
#include <deque>
using namespace std;

//Input: nums = [2,4,3,3,5,4,9,6], k = 4
//Output: [2,3,3,4]
//Input: nums = [3,5,2,6], k = 2
//Output: [2,6]
class Solution_1673 {
public:
    vector<int> mostCompetitive(vector<int>& nums, int k) {
        std::deque<int> q;
        int len = nums.size();
        for(int i = 0;i < len;i++){
            while(q.size() > 0 && q.back() > nums[i] && (k - q.size()) < (len - i)){
                q.pop_back();
            }
            if(q.size() < k){
                q.push_back(nums[i]);
            }
        }
        std::vector<int> res(k);
        int cap = q.size();
        for (int i = k - 1;i >= 0;i--){
            res[i] = q.back();
            q.pop_back();
        }
        return res;
    }
};