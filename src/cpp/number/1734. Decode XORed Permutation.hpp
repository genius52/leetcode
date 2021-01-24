#include <vector>
using namespace std;

//Input: encoded = [6,5,4,6]
//Output: [2,4,1,5,3]
class Solution_1734 {
public:
    vector<int> decode(vector<int>& encoded) {
        int len = encoded.size();
        int k = (len + 1) / 2;
        int total = 1;
        if((k|1) == k){
            total = 0;
        }
        std::vector<int> res(len + 1);
        int cur = 0;
        for(int i = 1;i < len;i += 2){
            cur ^= encoded[i];
        }
        int last = total ^ cur;
        res[0] = last;
        for(int i = 1;i <= len;i++){
            res[i] = encoded[i - 1] ^ last;
            last = res[i];
        }
        return res;
    }
};