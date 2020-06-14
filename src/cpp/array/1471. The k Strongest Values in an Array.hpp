#include <vector>
#include <algorithm>
using namespace std;

class Solution_1471 {
public:
    vector<int> getStrongest(vector<int>& arr, int k) {
        std::vector<int> res;
        int len = arr.size();
        if (len == 0)
            return res;
        std::sort(arr.begin(), arr.end());
        int mid = arr[(len - 1) / 2];
        int begin = 0;
        int end = len - 1;
        while (begin <= end && k > 0) {
            if (abs(arr[begin] - mid) > abs(arr[end] - mid)) {
                res.push_back(arr[begin]);
                begin++;
                k--;
            }
            else {
                res.push_back(arr[end]);
                end--;
                k--;
            }
        }
        return res;
    }
};