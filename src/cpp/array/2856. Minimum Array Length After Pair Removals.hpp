#include <vector>
#include <unordered_map>
#include <queue>
using namespace std;

class Solution {
public:
    int minLengthAfterRemovals(vector<int>& nums) {
        int len = nums.size();
        std::unordered_map<int,int> num_cnt;
        for(auto n : nums){
            num_cnt[n]++;
        }
        std::priority_queue<int, std::vector<int>, std::less<int> > big_top;
        for(auto key_val : num_cnt){
            //big_top.push({key_val.second,key_val.first})
            big_top.push(key_val.second);
        }
        int delete_cnt = 0;
        while (big_top.size() > 1){
            auto cnt1 = big_top.top();
            cnt1--;
            delete_cnt++;
            big_top.pop();
            auto cnt2  = big_top.top();
            cnt2--;
            delete_cnt++;
            big_top.pop();
            if (cnt1 > 0){
                big_top.push(cnt1);
            }
            if(cnt2 > 0){
                big_top.push(cnt2);
            }
        }
        return len - delete_cnt;
    }
};