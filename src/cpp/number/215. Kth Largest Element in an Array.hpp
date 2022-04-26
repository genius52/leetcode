#include <vector>
#include <deque>
using namespace std;

class Solution_215 {
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

    int findKthLargest2(vector<int>& nums, int k) {
        std::priority_queue<int, std::vector<int>, std::greater<int> > q;
        int i = 0;
        int len = nums.size();
        while(i < k){
            q.push(nums[i]);
            i++;
        }
        int res = 0;
        for(;i < len;i++){
            if(nums[i] <= q.top())
                continue;
            q.pop();
            q.push(nums[i]);
        }
        return q.top();
    }
};