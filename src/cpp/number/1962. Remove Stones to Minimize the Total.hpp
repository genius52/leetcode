#include <vector>
#include <queue>
using namespace std;

class Solution_1962 {
public:
    int minStoneSum(vector<int>& piles, int k) {
        std::priority_queue<int> q;
        int len = piles.size();
        int sum = 0;
        for (int i = 0; i < len; i++) {
            sum += piles[i];
            q.push(piles[i]);
        }
        while (k > 0) {
            int val1 = q.top();
            q.pop();
            sum -= (val1 / 2);
            q.push(val1 - val1 / 2);
            k--;
        }
        if (sum < 0) {
            return 0;
        }
        return sum;
    }
};