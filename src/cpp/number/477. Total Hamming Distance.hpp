#include <vector>
using namespace std;

class Solution_477 {
public:
    int totalHammingDistance(vector<int>& nums) {
        int len = nums.size();
        int res = 0;
        for(int k = 0;k < 31;k++){
            int tag = 1 << k;
            int one_cnt = 0;
            for(int i = 0;i < len;i++){
                if((nums[i]|tag) == nums[i]){
                    one_cnt++;
                }
            }
            res += (len - one_cnt) * one_cnt;
        }
        return res;
    }
};