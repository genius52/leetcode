#include <math.h>
#include <unordered_set>

class Solution_762 {
public:
    int countPrimeSetBits2(int L, int R) {
        int res = 0;
        std::unordered_set<int> s{2,3,5,7,11,13,17,19,23,29};
        for(int i = L;i <= R;i++){
            int bitcnt = 0;
            int n = i;
            while(n > 0){
                bitcnt += n & 1;
                n = n >> 1;
            }
            res += s.count(bitcnt);
        }
        return res;
    }

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