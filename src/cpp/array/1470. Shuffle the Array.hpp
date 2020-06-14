#include <vector>
using namespace std;

class Solution_1470 {
public:
    vector<int> shuffle(vector<int>& nums, int n) {
        std::vector<int> res;
        res.resize(n * 2);
        int i = 0;
        int pos = 0;
        while (i < n) {
            res[pos++] = nums[i];
            res[pos++] = nums[i + n];
            i++;
        }
        return res;
    }
};