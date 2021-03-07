#include <vector>
#include <map>
using namespace std;
 
//Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
//Output: [3,3,5,5,6,7]
class Solution_239 {
public:
    vector<int> maxSlidingWindow(vector<int>& nums, int k) {
        int len = nums.size();
        std::map<int,int,std::greater<int>> record;
        std::vector<int> res(len - k + 1);
        for(int i = 0;i < k;i++){
            record[nums[i]]++;
        }
        res[0] = record.begin()->first;
        for(int i = 1;i <= len - k;i++){
            record[nums[i - 1]]--;
            if(record[nums[i - 1]] == 0){
                record.erase(nums[i - 1]);
            }
            record[nums[i + k - 1]]++;
            res[i] = record.begin()->first;
        }
        return res;
    }
};