#include <vector>
#include <queue>
#include <set>
using namespace std;

//Given nums = [1,3,-1,-3,5,3,6,7], and k = 3.
//
//Window position                Median
//---------------               -----
//[1  3  -1] -3  5  3  6  7       1
// 1 [3  -1  -3] 5  3  6  7       -1
// 1  3 [-1  -3  5] 3  6  7       -1
// 1  3  -1 [-3  5  3] 6  7       3
// 1  3  -1  -3 [5  3  6] 7       5
// 1  3  -1  -3  5 [3  6  7]      6
class Solution_480 {
public:
    vector<double> medianSlidingWindow(vector<int>& nums, int k) {
        std::vector<double> res;
        int len = nums.size();
        if(len < k){
            return res;
        }
        int first = 0;
        int second = 0;
        if((k | 1) == k){
            first = k/2;
            second = first;
        }else{
            first = (k - 1)/2;
            second = first + 1;
        }
        std::multiset<int> s;
        for(int i = 0;i < k;i++){
            s.insert(nums[i]);
        }
        auto it1 = s.begin();
        std::advance(it1, first);
        long val1 = *it1;
        auto it2 = s.begin();
        std::advance(it2, second);
        long val2 = *it2;
        res.push_back((double(val1 + val2))/2);

        for(int i = k;i < len;i++){
            s.erase(s.find(nums[i - k]));
            s.insert(nums[i]);
            auto it1 = s.begin();
            std::advance(it1, first);
            long first_val = *it1;
            auto it2 = s.begin();
            std::advance(it2, second);
            long second_val = *it2;
            res.push_back((double(first_val + second_val))/2);
        }
        return res;
    }
};