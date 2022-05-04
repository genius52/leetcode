#include <vector>
using namespace std;

class Solution_581 {
public:
    int findUnsortedSubarray(vector<int>& nums) {
        vector<int> sort_arr;
        sort_arr.assign(nums.begin(), nums.end());
        std::sort(sort_arr.begin(),sort_arr.end());
        int diff = 0;
        int len = nums.size();
        int begin = -1;
        int end = -1;
        for(int i = 0;i < len;i++){
            if(nums[i] != sort_arr[i]){
                if(begin == -1){
                    begin = i;
                }
                end = i;
            }
        }
        if(begin == -1)
            return 0;
        return end - begin + 1;
    }
};