#include <vector>
#include <priority_queue>
#include <deque>
using namespace std;

class Solution {
public:
//    int findKthLargest(vector<int>& nums, int k) {
//        int len = nums.size();
//        std::nth_element(nums.begin(),nums.begin() + len - k,nums.end());
//        return nums[len - k];
//    }

    int findKthLargest(vector<int>& nums, int k) {
        std::priority_queue<int, std::deque<int>, std::less<int> > q(nums.begin(),nums.end());
        int res = 0;
        while(k > 0){
            res = q.top();
            q.pop();
            k--;
        }
        return res;
    }
};