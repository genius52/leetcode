#include <algorithm>
#include <numeric>
#include <vector>
#include <math.h>
using namespace std;

class Solution_1619 {
public:
    double trimMean(vector<int>& arr) {
        int l = arr.size();
        std::sort(arr.begin(), arr.end());
        int begin = l / 20;
        int end = l - l / 20;
        auto total = std::accumulate(arr.begin() + begin, arr.begin() + end,0);
        double a = (double)total / (double(l) * 0.9);
        double res = floor(a * 100000.000f + 0.5) / 100000.000f;
        return res;
    }
};