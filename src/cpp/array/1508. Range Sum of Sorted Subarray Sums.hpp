#include <vector>
#include <queue>
using namespace std;

//Input: nums = [1,2,3,4], n = 4, left = 3, right = 4
//Output: 6
//Explanation: The given array is the same as example 1. We have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10].
//The sum of the numbers from index le = 3 to ri = 4 is 3 + 3 = 6.
class Solution_1508 {
public:
    int rangeSum(vector<int>& nums, int n, int left, int right){
        std::vector<int> record(n);
        std::vector<int> pre_sum(n + 1);
        pre_sum[0] = nums[0];
        for(int i = 1;i <= n;i++){
            pre_sum[i] = nums[i - 1] + pre_sum[i - 1];
            pre_sum[i] %= 1000000007;
        }
        std::priority_queue<int,vector<int>,greater<int>> q;
        for(int i = 0;i <= n;i++){
            for(int j = i - 1;j >= 0;j--){
                q.push(pre_sum[i] - pre_sum[j]);
            }
        }
        int res = 0;
        int pos = 1;
        while(q.size() > 0){
            if(pos >= left && pos <= right){
                res += q.top();
                res %= 1000000007;
            }else if (pos > right){
                break;
            };
            q.pop();
            pos++;
        }
        return res;
    }
};