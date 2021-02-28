#include <vector>
#include <queue>
#include <algorithm>
#include <numeric>
using namespace std;


//Input: nums1 = [1, 2, 3, 4, 5, 6], nums2 = [1, 1, 2, 2, 2, 2]
//Output : 3
//Explanation : You can make the sums of nums1 and nums2 equal with 3 operations.All indices are 0 - indexed.
//- Change nums2[0] to 6. nums1 = [1, 2, 3, 4, 5, 6], nums2 = [6, 1, 2, 2, 2, 2].
//- Change nums1[5] to 1. nums1 = [1, 2, 3, 4, 5, 1], nums2 = [6, 1, 2, 2, 2, 2].
//- Change nums1[2] to 2. nums1 = [1, 2, 2, 4, 5, 1], nums2 = [6, 1, 2, 2, 2, 2].
class Solution_1775 {
public:
    int minOperations(vector<int>& nums1, vector<int>& nums2) {
        int len1 = nums1.size();
        int len2 = nums2.size();
        int sum1 = std::accumulate(nums1.begin(), nums1.end(), 0);
        int sum2 = std::accumulate(nums2.begin(), nums2.end(), 0);
        if (sum1 == sum2)
            return 0;
        std::priority_queue<int, std::deque<int>, std::greater<int> > smalltop_queue;
        std::priority_queue<int, std::deque<int>, std::less<int> > bigtop_queue;
        int big_sum = max(sum1, sum2);
        int small_sum = min(sum1, sum2);
        int short_len = min(len1, len2);
        int long_len = max(len1, len2);
        if (short_len * 6 < long_len)
            return -1;
        if (sum1 > sum2) {
            for (auto n : nums1) {
                bigtop_queue.push(n);
            }
            for (auto n : nums2) {
                smalltop_queue.push(n);
            }
        }
        else {
            for (auto n : nums1) {
                smalltop_queue.push(n);
            }
            for (auto n : nums2) {
                bigtop_queue.push(n);
            }
        }

        int steps = 1;
        while (small_sum != big_sum) {
            int diff = big_sum - small_sum;
            int small_top = smalltop_queue.top();
            int big_top = bigtop_queue.top();
            if (6 - small_top >= diff || big_top - 1 >= diff) {
                return steps;
            }
            if ((6 - small_top) > (big_top - 1)) {
                smalltop_queue.pop();
                smalltop_queue.push(6);
                small_sum += 6 - small_top;
            }
            else {
                bigtop_queue.pop();
                bigtop_queue.push(1);
                big_sum -= (big_top - 1);
            }
            steps++;
        }
        return steps;
    }
};