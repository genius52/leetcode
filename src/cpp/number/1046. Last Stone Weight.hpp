#include <set>
#include <vector>
using namespace std;

class Solution {
public:
    int lastStoneWeight(vector<int>& stones){
        int len = stones.size();
        std::priority_queue<int> q(stones.begin(),stones.end());
        while(q.size() > 1){
            auto top1 = q.top();
            q.pop();
            auto top2 = q.top();
            q.pop();
            auto diff = abs(top1 - top2);
            if(diff != 0)
                q.push(diff);
        }
        if(q.size() == 0)
            return 0;
        return q.top();
    }

    int lastStoneWeight2(std::vector<int>& stones) {
        std::multiset<int,std::greater<int>> s;
        for(auto it = stones.begin();it != stones.end();it++){
            s.insert(*it);
        }
        while(s.size()>0){
            auto first = s.begin();
            auto biggest = *first;
            s.erase(first++);
            if(first == s.end())
                return biggest;
            else{
                if(*first != biggest)
                    s.insert(biggest - *first);
                s.erase(first);
            }
        }
        return 0;
    }
};

