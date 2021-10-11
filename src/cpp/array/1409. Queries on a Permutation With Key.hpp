#include <vector>
using namespace std;

class Solution_1409 {
public:
    vector<int> processQueries(vector<int>& queries, int m) {
        vector<int> nums(m);
        for (int i = 1; i <= m; i++) {
            nums[i - 1] = i;
        }
        int len = queries.size();
        vector<int> res;
        for (int i = 0; i < len; i++) {
            int target = queries[i];
            int j = 0;
            for (; j <= m; j++) {
                if (nums[j] == target) {
                    break;
                }
            }
            int tmp = nums[j];
            res.push_back(j);
            for (int k = j; k > 0; k--) {
                nums[k] = nums[k - 1];
            }
            nums[0] = tmp;
        }
        return res;
    }
};