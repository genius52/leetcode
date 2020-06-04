#include <math.h>
#include <unordered_map>

class Solution_397 {
public:
    int dp_integerReplacement(long n,std::unordered_map<long,long>& record){
        if(n == 1)
            return 0;
        if(record.find(n) != record.end())
            return record[n];
        long min_cnt = 0;
        if(n % 2 == 0)
            min_cnt = 1 + dp_integerReplacement(n / 2,record);
        else
            min_cnt = min(1 + dp_integerReplacement(n + 1,record),1 + dp_integerReplacement(n - 1,record));
        record[n] = min_cnt;
        return min_cnt;
    }

    int integerReplacement(int n) {
        if(n <= 0)
            return 0;
        std::unordered_map<long,long> record;
        return dp_integerReplacement(n,record);
    }
};