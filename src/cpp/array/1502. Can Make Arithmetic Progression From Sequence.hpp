#include <vector>
#include <algorithm>
using namespace std;

class Solution_1502 {
public:
    bool canMakeArithmeticProgression(vector<int>& arr) {
        std::sort(arr.begin(),arr.end());
        int len = arr.size();
        if(len <= 1)
            return true;
        int diff = arr[1] - arr[0];
        for(int i = 2;i < len;i++){
            if(arr[i] - arr[i - 1] != diff)
                return false;
        }
        return true;
    }
};