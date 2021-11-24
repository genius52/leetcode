#include <vector>
#include <set>
using namespace std;

class Solution_1675 {
public:
    int minimumDeviation(vector<int>& nums) {
        std::set<int> s;
        for(auto n : nums){
            if((n | 1) == n){
                s.insert(n * 2);
            }else{
                s.insert(n);
            }
        }
        int res = *s.rbegin() - *s.begin();
        while ((*s.rbegin() | 1) != *s.rbegin()){
            auto big = *s.rbegin();
            s.erase(big);
            big /= 2;
            s.insert(big);
            res = min(res,*s.rbegin() - *s.begin());
        }
        return res;
    }
};