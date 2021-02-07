#include <vector>
using namespace std;

class Solution {
    std::vector<int> data;
    int len = 0;
public:
    Solution(vector<int>& nums) {
        data = nums;
        len = nums.size();
    }

    int pick(int target) {
        int find_cnt = 1;
        int res = 0;
        int i = 0;
        while(i < len){
            if(data[i] == target){
                if(rand() % find_cnt++ == 0)
                    res = i;
            }
            i++;
        }
        return res;
    }
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(nums);
 * int param_1 = obj->pick(target);
 */