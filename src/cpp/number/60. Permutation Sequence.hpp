#include <string>
#include <algorithm>
using namespace std;

class Solution_60 {
public:
    string getPermutation(int n, int k) {
        std::vector<int> data(n);
        int index = 0;
        for(int i = 1;i <= n;i++){
            data[index++] = i;
        }
        int step = 1;
        while(step < k){
            std::next_permutation(data.begin(),data.end());
            step++;
        }
        std::string res;
        for(int i = 0;i < n;i++){
            res += std::to_string(data[i]);
        }
        return res;
    }
};