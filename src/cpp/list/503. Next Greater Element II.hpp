#include <vector>
#include <stack>
using namespace std;

//Input: nums = [1,2,1]
//Output: [2,-1,2]
class Solution_503 {
public:
    vector<int> nextGreaterElements(vector<int>& nums) {
        int len = nums.size();
        std::stack<int> s;//store index with decrease order
        std::vector<int> res(len);
        for(int i = 0;i < len * 2;i++){
            int index = i;
            if(index >= len){
                index = index % len;
            }
            if(i < len)
                res[index] = -1;
            if(s.empty()){
                if(i < len)
                    s.push(index);
            }else{
                while(!s.empty() && nums[s.top()] < nums[index]){
                    int prev_index = s.top();
                    s.pop();
                    res[prev_index] = nums[index];
                }
                if(i < len)
                    s.push(index);
            }
        }
        return res;
    }
};