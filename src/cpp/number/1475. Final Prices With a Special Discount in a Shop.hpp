#include <vector>
#include <deque>
using namespace std;

//Input: prices = [8,4,6,2,3]
//Output: [4,2,4,2,3]
class Solution_1475 {
public:
    vector<int> finalPrices(vector<int>& prices) {
        int l = prices.size();
        std::vector<int> res = prices;
        if(l <= 1)
            return res;
        std::deque<int> q;
        q.push_back(0);
        for(int i = 1;i < l;i++){
            while(!q.empty() && prices[i] <= prices[q.back()]){
                res[q.back()] = res[q.back()] - prices[i];
                q.pop_back();
            }
            q.push_back(i);
        }
        return res;
    }
};