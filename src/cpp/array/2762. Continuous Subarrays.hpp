#include <vector>
#include <set>
using namespace std;

class Solution_2762 {
public:
    long long continuousSubarrays(vector<int>& nums) {
        int len = nums.size();
        std::multiset<int> data;
        long long res = 0;
        int left = 0;
        int right = 0;
        while (left < len){
            while(right < len && (data.empty() || ((abs(*data.rbegin() - nums[right]) <= 2) && (abs(*data.begin() - nums[right]) <= 2)))){
                data.insert(nums[right]);
                right++;
            }
            res += right - left;
            data.erase(data.find(nums[left]));
            left++;
        }
        return res;
    }
};