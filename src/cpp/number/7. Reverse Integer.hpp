#include <deque>
#include <cmath>

class Solution_7 {
public:
    int reverse(int x) {
        bool negative = x < 0;
        std::deque<int> s;
        while(x != 0){
            auto result = x % 10;
            s.push_back(result);
            x = x/10;
        }
        long val = 0;
        int cnt = 0;
        while(!s.empty()){
            long temp = s.back();
            s.pop_back();
            val += temp * pow(10,cnt++);
        }
        int int_max = pow(2,31)-1;
        int int_min = pow(-2,31);
        if(val > int_max)
            return 0;
        if(val < int_min)
            return 0;
        return val;
    }
};