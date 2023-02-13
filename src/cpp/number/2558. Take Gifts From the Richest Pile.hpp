#include <vector>
#include <queue>
#include <math.h>
using namespace std;

class Solution_2558 {
public:
    long long pickGifts(vector<int>& gifts, int k) {
        std::priority_queue<int, std::vector<int>, std::less<int> > big_top;
        long long sum = 0;
        for(auto g : gifts){
            big_top.push(g);
            sum += g;
        }
        while(k > 0){
            auto top = big_top.top();
            big_top.pop();
            int val = floor(sqrt(top));
            int diff = top - val;
            if(diff == 0){
                break;
            }
            sum -= diff;
            big_top.push(val);
            k--;
        }
        return sum;
    }
};