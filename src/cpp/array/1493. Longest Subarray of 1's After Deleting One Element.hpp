#include <vector>
using namespace std;

//Input: nums = [0,1,1,1,0,1,1,0,1]
//Output: 5
//Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
class Solution_1493 {
public:
    int calc_one_length(vector<int>& nums,int len,int pos,bool forward){
        if(pos >= len)
            return 0;
        int cnt = 0;
        while(pos < len && pos >= 0){
            if(nums[pos] == 1){
                if(forward)
                    pos++;
                else
                    pos--;
                cnt++;
            }else{
                break;
            }
        }
        return cnt;
    }

    int longestSubarray(vector<int>& nums) {
        int len = nums.size();
        int pos = 0;
        int max_len = 0;
        bool find_zero = false;
        while(pos < len){
            if(nums[pos] == 0){
                int pre = calc_one_length(nums,len,pos - 1,false);
                int post = calc_one_length(nums,len,pos + 1,true);
                if((pre + post) > max_len)
                    max_len = pre + post;
                if(post == 0)
                    pos++;
                else
                    pos += post;
                find_zero = true;
            }else{
                pos++;
            }
        }
        if(!find_zero)
            return len - 1;
        return max_len;
    }
};