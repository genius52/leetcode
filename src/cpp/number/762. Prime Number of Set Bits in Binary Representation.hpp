#include <math.h>

class Solution_762 {
public:
    int countPrimeSetBits(int L, int R) {
        int res = 0;
        for(int i = L;i <= R;i++){
            int bitcnt = count_bits(i);
            if(is_prime(bitcnt))
                res++;
        }
        return res;
    }

    int count_bits(int n){
        int bit = 1;
        int cnt = 0;
        while(bit < n){
            if(n == (n | bit))
                cnt++;
            bit = bit << 1;
        }
        return cnt;
    }

    bool is_prime(int n) {
        if (n <= 3) {
            return n > 1;
        }
        int m = static_cast<int>(sqrt(n));
        for (int i = 2; i <= m; i++) {
            if(n % i == 0) {
                return false;
            }
        }
        return true;
    }
};