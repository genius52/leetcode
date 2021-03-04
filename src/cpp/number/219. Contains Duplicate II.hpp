#include <vector>
#include <unordered_map>
#include <set>
using namespace std;

class Solution_219 {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        std::unordered_map<int,int> record;
        int len = nums.size();
        for(int i = 0;i < len;i++){
            if(record.find(nums[i]) != record.end()){
                if(abs(record[nums[i]] - i) <= k){
                    return true;
                }
            }
            record[nums[i]] = i;
        }
        return false;
    }
};