#include <cmath>

class Solution_9 {
public:
    bool isPalindrome(int x) {
        if(x < 0)
            return false;
        if(x == 0)
            return true;
        long new_val = 0;
        int n = 0;
        int origin_val = x;
        while(origin_val != 0)
        {
            int result = origin_val % 10;
            new_val = new_val * 10 + result;
            origin_val = origin_val / 10;
        }
        if(origin_val > pow(2,31) - 1)
            return false;
        return new_val == x;
    }
};