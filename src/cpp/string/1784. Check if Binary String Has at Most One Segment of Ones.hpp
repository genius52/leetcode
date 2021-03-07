#include <string>

class Solution_1784 {
public:
    bool checkOnesSegment(string s) {
        int len = s.size();
        int i = 0;
        while (i < len) {
            if (s[i] == '1') {
                i++;
                continue;
            }
            else {
                break;
            }
        }
        int j = len - 1;
        while (j >= 0) {
            if (s[j] == '0') {
                j--;
                continue;
            }
            else {
                break;
            }
        }
        i--;
        if (i >= j)
            return true;
        else
            return false;
    }
};
