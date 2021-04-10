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
        std::map<int,std::vector<int>> m;
        int len = Profits.size();
        for(int i = 0;i < len;i++){
            m[Capital[i]].push_back(Profits[i]);
        }
        std::priority_queue<int> q;
        std::map<int,std::vector<int>>::iterator it = m.begin();
        int cur_capital = W;
        for(int i = 0;i < k;i++){
            for(;it != m.end();it++) {
                if(cur_capital < it->first)
                    break;
                vector<int> tmp = it->second;
                for(auto it2 = tmp.begin();it2 != tmp.end();it2++) {
                    q.push(*it2);
                }
            }
            if(q.empty())
                break;
            cur_capital += q.top();
            q.pop();
        }
        return cur_capital;
    }
};