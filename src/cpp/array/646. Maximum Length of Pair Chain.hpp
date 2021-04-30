#include <vector>
using namespace std;

class Solution {
public:
    int findLongestChain(vector<vector<int>>& pairs) {
        std::sort(pairs.begin(),pairs.end(),[](const std::vector<int>& p1, const std::vector<int>& p2){
            return p1[1] < p2[1];
        });
        int cnt = 0;
        int pre_end = -2147483648;
        for(auto p : pairs){
            if(p[0] > pre_end){
                cnt++;
                pre_end = p[1];
            }
        }
        return cnt;
    }
};