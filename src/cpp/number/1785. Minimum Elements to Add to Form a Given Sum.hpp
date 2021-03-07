#include <vector>
using namespace std;

class Solution_1785 {
public:
    int minElements(vector<int>& nums, int limit, int goal) {
        long sum = 0;
        int len = nums.size();
        for (int i = 0; i < len; i++) {
            sum += nums[i];
        }
        long need = long(abs(long(goal) - sum));
        if (need == 0)
            return 0;
        long cnt = 0;
        cnt = need / limit;
        if ((need % limit) != 0)
            cnt++;
        return cnt;
    }
};