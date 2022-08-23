#include <vector>
#include <set>
using namespace std;

class Solution_2382 {
public:
    vector<long long> maximumSegmentSum(vector<int>& nums, vector<int>& removeQueries) {
        int len = nums.size();
        std::vector<long long> prefix_sum(len + 1);
        for(int i = 0;i < len;i++){
            prefix_sum[i + 1] = prefix_sum[i] + nums[i];
        }
        std::set<long> delete_idx;//删除的位置
        delete_idx.insert(0);
        delete_idx.insert(len + 1);
        std::multiset<long long> sum;
        sum.insert(prefix_sum[len]);
        int len2 = removeQueries.size();
        std::vector<long long> res(len2);
        for(int i = 0;i < len2;i++){
            int pos = removeQueries[i] + 1;
            auto find = delete_idx.upper_bound(pos);
            int left = *std::prev(find);
            int right = *find;
            long long last_sum = prefix_sum[right - 1] - prefix_sum[left];
            sum.erase(sum.find((last_sum)));
            if (pos - left - 1 > 0){
                auto left_sum = prefix_sum[pos - 1] - prefix_sum[left];
                sum.insert(left_sum);
            }
            if(right - pos - 1 > 0){
                auto right_sum = prefix_sum[right - 1] - prefix_sum[pos];
                sum.insert(right_sum);
            }
            delete_idx.insert(pos);
            if(sum.empty())
                res[i] = 0;
            else
                res[i] = *sum.rbegin();
        }
        return res;
    }
};