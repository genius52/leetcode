#include <vector>
using namespace std;

//[1,2,10,5,7]
class Solution_1909 {
public:
    bool canBeIncreasing(vector<int>& nums) {
        int len = nums.size();
        bool del = false;
        int pre_idx = 0;
        for(int i = 1;i < len;i++){
            if(nums[pre_idx] < nums[i]){
                pre_idx = i;
                continue;
            }
            if(del){
                return false;
            }
            del = true;
            if((i - 1) > 0){
                if(nums[i - 2] < nums[i]){
                    pre_idx = i;
                }
            }else{
                pre_idx = i;
            }
        }
        return true;
    }
};