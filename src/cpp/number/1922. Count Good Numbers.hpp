#include <map>

class Solution_1922 {
public:
    int countGoodNumbers(long long n) {
        long long i = 0;
        std::map<long long, long long> record;
        record[1] = 5;
        record[2] = 20;
        long long res = 1;
        while (i < n) {
            long long rest = n - i;
            if (rest == 1) {
                res *= 5;
                res %= 1000000007;
                break;
            }
            auto find = record.upper_bound(rest);
            find--;
            res = res * (*find).second % 1000000007;
            i += (*find).first;
            record[i] = res;
        }
        return (int)res;
    }
};