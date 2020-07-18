#include <unordered_map>
using namespace std;

//Input: 10
//Output: 4
//Explanation:
//There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
class Solution_788 {
public:
    std::unordered_map<int,int> valid_rotate;
    bool check_rotate(int n){
        int new_num = 0;
        int old_num = n;
        int upgrade = 1;
        while(n > 0){
            int rest = n % 10;
            if(valid_rotate.find(rest) == valid_rotate.end())
                return false;
            new_num += valid_rotate[rest] * upgrade;
            upgrade *= 10;
            n = n / 10;
        }
        return !(new_num == old_num);
    }

    int rotatedDigits(int N) {
        //2 <-> 5 6 <-> 9
        //0 <-> 0 1 <-> 1 2<->5
        valid_rotate[0] = 0;
        valid_rotate[1] = 1;
        valid_rotate[2] = 5;
        valid_rotate[5] = 2;
        valid_rotate[6] = 9;
        valid_rotate[8] = 8;
        valid_rotate[9] = 6;
        int res = 0;
        for(int i = 2;i <= N;i++){
            if(check_rotate(i))
                res++;
        }
        return res;
    }
};