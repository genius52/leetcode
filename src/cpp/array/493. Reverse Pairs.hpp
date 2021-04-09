#include <vector>
#include <set>
using namespace std;

//Input: [2,4,3,5,1]
//Output: 3
class Solution_493 {
public:
    int merge_sort(std::vector<int>& nums,int begin,int end){
        if(end - begin <= 1){
            return 0;
        }
        int mid = (begin + end)/2;
        int count = merge_sort(nums,begin,mid) + merge_sort(nums,mid,end);
        int j = mid;
        for(int i = begin;i <= mid;i++){
            while(j < end && long(nums[i]) > (long(nums[j]) * 2)){
                j++;
            }
            count += j - mid;
        }
        std::inplace_merge(nums.begin() + begin,nums.begin() + mid,nums.begin() + end);
        return count;
    }

    int reversePairs(vector<int>& nums){
        int len = nums.size();
        return merge_sort(nums,0,len - 1);
    }

//binary search
//    int reversePairs(vector<int>& nums){
//        int cnt = 0;
//        int len = nums.size();
//        std::multiset<int> s;
//        for(int i = 0;i < len;i++){
//            auto it = s.upper_bound(nums[i] * 2);
//            if(it != s.end()){
//                cnt += std::distance(it, s.end());
//            }
//            s.insert(nums[i]);
//        }
//        return cnt;
//    }

    //brute force
//    int reversePairs(vector<int>& nums) {
//        int cnt = 0;
//        int len = nums.size();
//        for(int i = 0;i < len;i++){
//            for(int j = i + 1;j < len;j++){
//                long right = 2 * long(nums[j]);
//                if(nums[i] > right){
//                    cnt++;
//                }
//            }
//        }
//        return cnt;
//    }
};