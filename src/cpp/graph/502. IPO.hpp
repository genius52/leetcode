#include <vector>
#include <map>
#include <set>
#include <unordered_set>
#include <queue>
using namespace std;

//Input: k=2, W=0, Profits=[1,2,3], Capital=[0,1,1].
//
//Output: 4
class Solution_502 {
public:
    int findMaximizedCapital(int k, int W, vector<int>& Profits, vector<int>& Capital) {
        std::map<int,std::multiset<int,std::greater<int>>> m;
        int len = Profits.size();
        for(int i = 0;i < len;i++){
            m[Capital[i]].insert(Profits[i]);
        }
        int cur_capital = W;
        for(int i = 0;i < k;i++){
            int cur_max_pro = 0;
            int key = -1;
            std::multiset<int>::iterator itval;
            for(auto it = m.begin();it != m.end();it++) {
                if(cur_capital < it->first)
                    break;
                int pro = *it->second.begin();
                if(pro > cur_max_pro) {
                    cur_max_pro = pro;
                    key = it->first;
                    itval = it->second.begin();
                }
            }
            if(key != -1){
                if(m[key].size() == 1){
                    m.erase(key);
                }else{
                    m[key].erase(itval);
                }
            }else{
                break;
            }
            cur_capital += cur_max_pro;
        }
        return cur_capital;
    }
};