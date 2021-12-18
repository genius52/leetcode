#include <vector>
#include <set>
#include <math.h>
using namespace std;

class Solution_1818 {
public:
    int minAbsoluteSumDiff(vector<int>& nums1, vector<int>& nums2){
        std::vector<int> sorted = nums1;
        std::sort(sorted.begin(),sorted.end());
        int len = nums1.size();
        long sum = 0;
        int max_decrease_diff = 0;
        for(int i = 0;i < len;i++){
            int origin_diff = abs(nums1[i] - nums2[i]);
            if(origin_diff == 0)
                continue;
            sum += origin_diff;
            int left = 0;
            int right = len - 1;
            while(left < right){
                int mid = (left + right)/2;
                if(sorted[mid] < nums2[i]){
                    left = mid + 1;
                }else if(sorted[mid] > nums2[i]){
                    right = mid;
                }else{
                    left = mid;
                    break;
                }
            }
            int diff = abs(nums2[i] - sorted[left]);
            if(left < len - 1){
                diff = min(diff, abs(nums2[i] - sorted[left + 1]));
            }
            if(left > 0){
                diff = min(diff, abs(nums2[i] - sorted[left - 1]));
            }
            max_decrease_diff = max(max_decrease_diff,origin_diff - diff);
        }
        return (sum - max_decrease_diff) % int(1e9 + 7);
    }
//    int minAbsoluteSumDiff(vector<int>& nums1, vector<int>& nums2) {
//        int len = nums1.size();
//        int sum = 0;
//        int max_gap = 0;
//        std::set<int> s1;
//        for(int i = 0;i < len;i++){
//            s1.insert(nums1[i]);
//            int gap = abs(nums1[i] - nums2[i]);
//            sum += gap;
//            sum = sum % 1000000007;
//        }
//        if(sum == 0)
//            return 0;
//        int minus_gap = 0;
//        for(int i = 0;i < len;i++){
//            int pre_gap = abs(nums1[i] - nums2[i]);
//            if(pre_gap == 0)
//                continue;
//            if(pre_gap < minus_gap)
//                continue;
//            int target1 = nums2[i];//在s1中寻找合适的替换数字
//
//            auto it1 = std::lower_bound(s1.begin(),s1.end(),target1);
//            int gap1 = 2147483647;
//            if(it1 != s1.end()){
//                gap1 = abs(*it1 - nums2[i]);
//            }
//
//            auto it3 = std::upper_bound(s1.begin(),s1.end(),target1);
//            int gap3 = 2147483647;
//            if(it3 != s1.begin()){
//                it3--;
//                gap3 = abs(*it3 - nums2[i]);
//            }
//            int new_gap1 = min(gap1,gap3);
//            if(new_gap1 < pre_gap){
//                int diff = pre_gap - new_gap1;
//                if(diff > minus_gap){
//                    minus_gap = diff;
//                }
//            }
//        }
//        return sum - minus_gap;
//    }
};