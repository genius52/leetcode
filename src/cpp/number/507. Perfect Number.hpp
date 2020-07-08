#include <math.h>

class Solution_507 {
public:
    bool checkPerfectNumber(int num) {
        if(num == 1)
            return false;
        if(num <= 0)
            return false;
        int limit = static_cast<int>(sqrt(num));
        int total = 1;
        for(int i = 2;i <= limit;i++){
            int rest = num % i;
            if(rest == 0){
                total += i;
                total += num / i;
            }
        }
        return total == num;
    }
};