#include <vector>
#include <queue>
#include <map>
using namespace std;

//Input: [-1, 3, 2, 0]
//
//Output: True
//
//Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].
class Solution_456 {
public:
    bool find132pattern(vector<int>& nums) {
        int len = nums.size();
        if(len < 3)
            return false;
        priority_queue<int,vector<int>,greater<int>> q;
        std::map<int,int> right_map;
        for(int i = 1;i < len;i++){
            right_map[nums[i]]++;
        }
        for(int i = 1;i < len - 1;i++){
            q.push(nums[i - 1]);
            auto cnt = right_map[nums[i]];
            if(cnt == 1)
                right_map.erase(nums[i]);
            else
                right_map[nums[i]]--;
            if(q.top() >= nums[i])
                continue;
            auto smallest = q.top();
            auto low = right_map.upper_bound(smallest);
            if(low != right_map.end()){
                int mid = (*low).first;
                if(mid > smallest && mid < nums[i])
                    return true;
            }
//            for(int k = i + 1;k < len;k++){
//                if(nums[k] < nums[i]){
//                    if(nums[k] > q.top())
//                        return true;
//                }
//            }
        }
        return false;
    }
};