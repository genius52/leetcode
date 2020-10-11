#include <string>
using namespace std;

class Solution_5535 {
public:
    int maxDepth(string s) {
        int left_cnt = 0;
        int right_cnt = 0;
        int len = s.size();
        int res = 0;
        int pos = 0;
        while (pos < len)
        {
            if (s[pos] == '(') {
                left_cnt++;
                if (left_cnt > res)
                    res = left_cnt;
            }
            else if (s[pos] == ')') {
                left_cnt--;
            }
            pos++;
        }
        return res;
    }
};