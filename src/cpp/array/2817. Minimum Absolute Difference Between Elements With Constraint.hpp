#include <vector>
#include <set>
using namespace std;

class Solution_2817 {
public:
    int minAbsoluteDifference(vector<int>& nums, int x) {
        if(x == 0){
            return 0;
        }
        int len = nums.size();
        std::multiset<int> s;
        int res = 1 << 31 - 1;
        for(int i = x;i < len;i++){
            s.insert(nums[i - x]);
            s.insert(nums[i]);
            auto find = s.find(nums[i]);
            auto left = find;
            auto right = find;
            if(left != s.begin()){
                left--;
                res = min(res, abs(*left - nums[i]));
            }
            right++;
            if(right != s.end()){
                res = min(res, abs(*right - nums[i]));
            }
            s.erase(find);
        }
        return res;
    }
};