#include <vector>
#include <unordered_map>
#include <map>
using namespace std;

class FindSumPairs {
    std::vector<int> nums1_;
    std::unordered_map<int,int> map2_;
    std::vector<int> nums2_;
public:
    FindSumPairs(vector<int>& nums1, vector<int>& nums2) {
        nums1_ = nums1;
        std::sort(nums1_.begin(),nums1_.end());
        for(auto n : nums2){
            map2_[n]++;
        }
        nums2_ = nums2;
    }

    void add(int index, int val) {
        int old_val = nums2_[index];
        int new_val = nums2_[index] + val;
        nums2_[index] += val;
        map2_[old_val]--;
        if(map2_[old_val] == 0){
            map2_.erase(old_val);
        }
        map2_[new_val]++;
    }

    int count(int tot) {
        int res = 0;
        for(auto it1 = nums1_.begin();it1 != nums1_.end() && *it1 <= tot;it1++){
            int target = tot - *it1;
            if(map2_.find(target) != map2_.end()){
                res += map2_[target];
            }
        }
        return res;
    }
};