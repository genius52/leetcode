#include <vector>
#include <unordered_map>
#include <math.h>
using namespace std;

class Solution_952 {
public:
    int find_root(std::vector<int>& groups,int i){
        if(groups[i] != i){
            groups[i] = find_root(groups,groups[i]);
        }
        return groups[i];
    }
    int largestComponentSize(vector<int>& nums) {
        int len = nums.size();
        int max_num = 0;
        for(auto n : nums){
            if(n > max_num){
                max_num = n;
            }
        }
        std::vector<int> groups(max_num + 1);
        for(int i = 0;i <= max_num;i++){
            groups[i] = i;
        }
        for(int i = 0;i < len;i++){
            int limit = sqrt(nums[i]);
            for(int j = 2;j <= limit;j++){
                if (nums[i] % j == 0){
                    int mod = nums[i] / j;
                    groups[find_root(groups,nums[i])] = groups[find_root(groups,j)];
                    groups[find_root(groups,nums[i])] = groups[find_root(groups,mod)];
                }
            }
        }
        std::unordered_map<int,int> group_cnt;
        int res = 0;
        for(int i = 0;i < len;i++){
            int group_id = find_root(groups,nums[i]);
            group_cnt[group_id]++;
            if(group_cnt[group_id] > res){
                res = group_cnt[group_id];
            }
        }
        return res;
    }
};