#include <vector>
#include <algorithm>
using namespace std;

class Solution_31 {
public:
    void nextPermutation(vector<int>& nums) {
        std::next_permutation(nums.begin(),nums.end());
    }
};