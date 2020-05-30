#include <vector>
using namespace std;

//Input: [1,17,5,10,13,15,10,5,16,8]
//Output: 7
//Explanation: There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].
class Solution_376 {
public:
    int dp_wiggleMaxLength(vector<int>& nums,int cur_pos,int pre_num,int pre_length,bool increase){
        if(cur_pos >= nums.size())
            return pre_length;
        int continue_pre = 0;
        int start_from_pre = 0;
        int start_from_current = 0;
        if((nums[cur_pos] > pre_num && increase) || (nums[cur_pos] < pre_num && !increase)){
            continue_pre = max(dp_wiggleMaxLength(nums,cur_pos + 1,nums[cur_pos],pre_length + 1,!increase),
                       dp_wiggleMaxLength(nums,cur_pos + 1,pre_num,pre_length,!increase));
        }
        else{
            continue_pre = dp_wiggleMaxLength(nums,cur_pos + 1,pre_num,pre_length,increase);
            start_from_pre = max(dp_wiggleMaxLength(nums,cur_pos + 1,pre_num,1,true),
                    dp_wiggleMaxLength(nums,cur_pos + 1,pre_num,pre_length,false));
        }
        //restart from current
        start_from_current = max(dp_wiggleMaxLength(nums,cur_pos + 1,nums[cur_pos],1,true),
                                     dp_wiggleMaxLength(nums,cur_pos + 1,nums[cur_pos],1,false));
        return max(max(pre_length,continue_pre),max(start_from_pre,start_from_current));
    }

//    int wiggleMaxLength(vector<int>& nums) {
//        return max(dp_wiggleMaxLength(nums,1,nums[0],1,true),dp_wiggleMaxLength(nums,1,nums[0],1,false));
//    }

    int wiggleMaxLength(vector<int>& nums) {
        int len = nums.size();
        if(len <= 1)
            return len;
//        std::vector<int> up(len,0);
//        std::vector<int> down(len,0);
        int lastup = 1;
        int lastdown = 1;
//        int currentup = 1;
//        int currentdown = 1;
        for(int i = 1;i < len;i++){
            if(nums[i] > nums[i - 1]){
                lastup = lastdown + 1;
            }else if(nums[i] < nums[i - 1]){
                lastdown = lastup + 1;
            }
        }
        return max(lastup,lastdown);
    }
};