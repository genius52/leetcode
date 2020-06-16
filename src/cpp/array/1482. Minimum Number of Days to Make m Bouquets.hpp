#include <vector>
#include <set>

class Solution_1482 {
public:
    int minDays(vector<int>& bloomDay, int m, int k) {
        int len = bloomDay.size();
        if (m * k > len)
            return -1;
        std::set<int> s(bloomDay.begin(), bloomDay.end(), std::less<int>());
        std::set<int>::reverse_iterator it = s.rbegin();
        int min_days = *it;
        while (it != s.rend()) {
            int flowers = 0;
            int i = 0;
            int days = *it;
            while (i <= len - k) {
                bool can_bloom = true;
                for (int j = 0; j < k; j++) {
                    if (bloomDay[i + j] > days) {
                        can_bloom = false;
                        i = i + j + 1;
                        break;
                    }
                }
                if (can_bloom) {
                    flowers++;
                    i = i + k;
                }
                if (flowers >= m)
                    break;
            }
            if (flowers >= m)
                min_days = *it;
            else
                return min_days;
            it++;
        }
        return min_days;
    }
};