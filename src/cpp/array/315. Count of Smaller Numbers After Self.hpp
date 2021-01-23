#include <vector>
#include <map>
#include <set>
using namespace std;

//Input: nums = [5,2,6,1]
//Output: [2,1,1,0]
//Explanation:
//To the right of 5 there are 2 smaller elements (2 and 1).
//To the right of 2 there is only 1 smaller element (1).
//To the right of 6 there is 1 smaller element (1).
//To the right of 1 there is 0 smaller element.
class Solution_315 {
public:
    vector<int> countSmaller(vector<int>& nums) {
        int len = nums.size();
        if(len == 0){
            return std::vector<int>();
        }else if(len == 1){
            return std::vector<int>{0};
        }
        std::vector<int> res(len);
        std::vector<int> sorted_data;
        res[len - 1] = 0;
        sorted_data.push_back(nums[len - 1]);
        for(int i = len - 2;i >= 0;i--){
            int cur = nums[i];
            auto it = std::lower_bound(sorted_data.begin(),sorted_data.end(),cur);
            res[i] = it - sorted_data.begin();
            sorted_data.insert(it,cur);
        }
        return res;
    }
};