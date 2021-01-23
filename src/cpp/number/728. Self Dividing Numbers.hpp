#include <vector>
using namespace std;

class Solution_728 {
public:
    vector<int> selfDividingNumbers(int left, int right) {
        vector<int> v;
        if(left > right)
            return v;
        for(int i = left;i <= right;i++){
            if(judge(i))
                v.push_back(i);
        }
        return v;
    }

    bool judge(int i){
        int origin = i;
        while(i > 0){
            int mod = i % 10;
            if(mod == 0)
                return false;
            if(origin % mod != 0)
                return false;
            i = i/10;
        }
        return true;
    }
};