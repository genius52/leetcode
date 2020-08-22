#include <vector>
using namespace std;

//Input: [1,4], 2
//Output: 4
class Solution_495 {
public:
    int findPoisonedDuration(vector<int>& timeSeries, int duration) {
        int len = timeSeries.size();
        if (len == 0)
            return 0;
        int cur_time = timeSeries[0] + duration;
        int total = duration;
        for (int i = 1; i < len; i++) {
            if (timeSeries[i] < cur_time) {
                total = total + timeSeries[i] - timeSeries[i - 1];
            }
            else {
                total += duration;
            }
            cur_time = timeSeries[i] + duration;
        }
        return total;
    }
};