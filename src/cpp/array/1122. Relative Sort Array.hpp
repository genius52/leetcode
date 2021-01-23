#include <vector>
#include <algorithm>
using namespace std;

class Solution {
public:
    std::vector<int> relativeSortArray(std::vector<int>& A, std::vector<int>& B) {
        int key[1001], idx = 0;
        for (int i = 0; i < 1001; i++) key[i] = 1000+i;
        for (int b : B) key[b] = idx++;
        sort(A.begin(), A.end(), [&](int c, int d){
            return key[c] < key[d];
        });
        return A;
    }
};
