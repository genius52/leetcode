#include <vector>
#include <queue>
using namespace std;

class Solution_2163 {
public:
    long long minimumDifference(vector<int>& nums) {
        std::priority_queue<long long, std::vector<long long>, std::greater<long long> > big_nums;//保存后半部分
        std::priority_queue<long long, std::vector<long long>, std::less<long long> > small_nums;//保存前半部分
        int total = nums.size();
        int l = total/3;
        long long prefix_sum = 0;
        long long suffix_sum = 0;
        for(int i = 0;i < l;i++){
            prefix_sum += nums[i];
            small_nums.push(nums[i]);
        };
        for(int i = total - 1;i >= l * 2;i--){
            suffix_sum += nums[i];
            big_nums.push(nums[i]);
        }
        //0...l - 1,l...2l - 1,2l ...3l - 1
        std::vector<long long> dp_prefix(l * 3 + 1);
        std::vector<long long> dp_suffix(l * 3 + 1);
        dp_prefix[l - 1] = prefix_sum;
        dp_suffix[2 * l] = suffix_sum;
        long long res = prefix_sum - suffix_sum;
        for(int i = l;i < l * 2;i++){
            prefix_sum += nums[i];
            small_nums.push(nums[i]);
            prefix_sum -= small_nums.top();
            small_nums.pop();
            dp_prefix[i] = prefix_sum;
        }
        for(int i = 2 * l - 1;i >= l;i--){
            suffix_sum += nums[i];
            big_nums.push(nums[i]);
            suffix_sum -= big_nums.top();
            big_nums.pop();
            dp_suffix[i] = suffix_sum;
        }

        for(int i = l - 1;i < 2 * l;i++){
            long long diff = dp_prefix[i] - dp_suffix[i + 1];
            if(diff < res){
                res = diff;
            }
        }
        return res;
    }
};