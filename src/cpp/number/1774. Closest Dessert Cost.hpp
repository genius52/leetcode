#include <vector>
#include <math.h>
using namespace std;
//Input: baseCosts = [2,3], toppingCosts = [4,5,100], target = 18
//Output: 17
//Explanation: Consider the following combination (all 0-indexed):
//- Choose base 1: cost 3
//- Take 1 of topping 0: cost 1 x 4 = 4
//- Take 2 of topping 1: cost 2 x 5 = 10
//- Take 0 of topping 2: cost 0 x 100 = 0
class Solution_1774 {
public:
    int min_diff = 2147483647;
    int dp_closestCost(vector<int>& toppingCosts, int top_len, int top_index, int cur_val, int target) {
        if (top_index >= top_len) {
            return cur_val;
        }
        if (cur_val >= target) {
            return cur_val;
        }
        else {
            int choose_one = 0;
            int choose_two = 0;
            int skip_cur = 0;
            choose_one = dp_closestCost(toppingCosts, top_len, top_index + 1, cur_val + toppingCosts[top_index], target);
            choose_two = dp_closestCost(toppingCosts, top_len, top_index + 1, cur_val + toppingCosts[top_index] * 2, target);
            skip_cur = dp_closestCost(toppingCosts, top_len, top_index + 1, cur_val, target);
            int diff_choose_one = abs(choose_one - target);
            int diff_choose_two = abs(choose_two - target);
            int diff_skip = abs(skip_cur - target);
            int cur_mindiff = min(diff_skip, min(diff_choose_one, diff_choose_two));
            if (choose_one == (target - cur_mindiff) || choose_two == (target - cur_mindiff) || skip_cur == (target - cur_mindiff)) {
                return target - cur_mindiff;
            }
            else {
                return target + cur_mindiff;
            }
        }
    }

    int closestCost(vector<int>& baseCosts, vector<int>& toppingCosts, int target) {
        int len1 = baseCosts.size();
        int len2 = toppingCosts.size();
        std::sort(baseCosts.begin(), baseCosts.end());
        std::sort(toppingCosts.begin(), toppingCosts.end());
        int res = 0;
        int min_diff = 2147483647;
        for (int i = 0; i < len1; i++) {
            if (baseCosts[i] >= target) {
                int diff = abs(target - baseCosts[i]);
                if (diff < min_diff) {
                    min_diff = diff;
                    return baseCosts[i];
                }
                break;
            }

            int diff = abs(target - baseCosts[i]);
            if (diff < min_diff) {
                min_diff = diff;
                res = target - diff;
            }
            int cur_target = dp_closestCost(toppingCosts, len2, 0, baseCosts[i], target);
            int cur_diff = abs(target - cur_target);
            if (cur_diff < min_diff) {
                min_diff = cur_diff;
                res = cur_target;
            }
            else if (cur_diff == min_diff) {
                if (cur_target < res) {
                    res = cur_target;
                }
            }
        }
        return res;
    }
};