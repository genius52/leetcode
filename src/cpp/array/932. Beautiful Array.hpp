#include <algorithm>
#include <vector>
using namespace std;

class Solution_932 {
public:
    vector<int> beautifulArray(int N) {
        std::vector<int> res;
        res.resize(N);
        for(int i = 0;i < N;i++){
            res[i] = i + 1;
        }
        for(int i = 0;i < N;i++){
            std::next_permutation(res.begin(),res.end());
        }
        return res;
    }
};