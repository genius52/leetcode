#include <vector>
#include <unordered_map>

//Input: [1,1,1,3,3,2,2,2]
//Output: [1,2]
class Solution_229 {
public:
    std::vector<int> majorityElement(std::vector<int>& nums) {
        std::unordered_map<int,int> record;
        std::vector<int> res;
        for (int i = 0;i < nums.size();i++){
            if (record.size() < 3){
                record[nums[i]] = 1;
            }else{
                
            }
        }
        return res;
    }
};