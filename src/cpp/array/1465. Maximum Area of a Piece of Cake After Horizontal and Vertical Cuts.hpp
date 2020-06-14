#include <vector>
using namespace std;

class Solution_1465 {
public:
    int maxArea(int h, int w, vector<int>& horizontalCuts, vector<int>& verticalCuts) {
        int res = 0;
        std::sort(horizontalCuts.begin(), horizontalCuts.end());
        std::sort(verticalCuts.begin(), verticalCuts.end());
        int max_rowspan = 0;
        int len_hor = horizontalCuts.size();
        for (int i = 1; i < len_hor; i++) {
            if ((horizontalCuts[i] - horizontalCuts[i - 1]) > max_rowspan) {
                max_rowspan = horizontalCuts[i] - horizontalCuts[i - 1];
            }
        }
        if (horizontalCuts[0] > max_rowspan) {
            max_rowspan = horizontalCuts[0];
        }
        if ((h - horizontalCuts[len_hor - 1]) > max_rowspan) {
            max_rowspan = h - horizontalCuts[len_hor - 1];
        }
        int len_ver = verticalCuts.size();
        int max_colspan = 0;
        for (int j = 1; j < len_ver; j++) {
            if ((verticalCuts[j] - verticalCuts[j - 1]) > max_colspan) {
                max_colspan = verticalCuts[j] - verticalCuts[j - 1];
            }
        }
        if (verticalCuts[0] > max_colspan) {
            max_colspan = verticalCuts[0];
        }
        if ((w - verticalCuts[len_ver - 1]) > max_colspan) {
            max_colspan = w - verticalCuts[len_ver - 1];
        }
        return long(max_rowspan) * long(max_colspan) % 1000000007;
    }
};
