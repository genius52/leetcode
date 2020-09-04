#include <vector>
#include <unordered_map>
#include <queue>
#include <math.h>
using namespace std;

class Solution_525 {
public:
    int findMaxLength(vector<int>& nums) {
        std::unordered_map<int,int> record;
        int max_len = 0;
        int zero_cnt = 0;
        int len = nums.size();
        for(int i = 0;i < len;i++){
            if(nums[i] == 0)
                zero_cnt++;
            int diff = i + 1 - zero_cnt - zero_cnt;
            if(diff == 0)
                max_len = i + 1;
            else{
                if(record.find(diff) != record.end()){
                    max_len = max(max_len,i - record[diff]);
                }else{
                    record[diff] = i;
                }
            }
        }
        return max_len;
    }
};