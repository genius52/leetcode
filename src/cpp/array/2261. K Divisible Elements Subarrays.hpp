#include <vector>
#include <unordered_set>
#include <set>
using namespace std;

class Solution_2261 {
public:
    int countDistinct(vector<int>& nums, int k, int p) {
        auto my_hash = [](const std::vector<int>& v) {
            long res = 0;
            int l = v.size();
            int MOD = 1e9 + 7;
            for(int i = 0;i < l;i++){
                res *= 200;
                res += v[i];
                res %= MOD;
            }
            return res;
        };
        auto my_equal = [](const std::vector<int>& v1, const std::vector<int>& v2) {
            int l1 = v1.size();
            int l2 = v2.size();
            if(l1 != l2){
                return false;
            }
            for(int i = 0;i < l1;i++){
                if(v1[i] != v2[i]){
                    return false;
                }
            }
            return true;
        };
        std::unordered_set<std::vector<int>, decltype(my_hash), decltype(my_equal)> res(10001, my_hash, my_equal);
        //std::unordered_set<std::vector<int>> res;
        int len = nums.size();
        for(int i = 0;i < len;i++){
            std::vector<int> cur;
            int cnt = 0;
            for(int j = i;j < len;j++){
                if(nums[j] % p == 0){
                    cnt++;
                }
                if(cnt > k){
                    break;
                }
                cur.push_back(nums[j]);
                res.insert(cur);
            }
        }
        return res.size();
    }
};