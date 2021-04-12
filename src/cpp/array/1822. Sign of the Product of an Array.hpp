#include <vector>
using namespace std;

class Solution_1822 {
public:
    int arraySign(vector<int>& nums) {
        int len = nums.size();
        int res = 1;
        for(int i = 0;i < len;i++){
            if(nums[i] == 0){
                return 0;
            }
            if(nums[i] < 0){
                res = -res;
            }
        }
        return res;
    }
};