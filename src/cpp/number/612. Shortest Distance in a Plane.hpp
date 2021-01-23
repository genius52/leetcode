#include <vector>
#include <queue>
using namespace std;

class Solution_612 {
public:
    int leastInterval(std::vector<char>& tasks, int n) {
        std::priority_queue<int>  q;
        std::map<char,int> record;

        for(auto t : tasks){
            if (record.find(t) == record.end()){
                record[t] = 1;
            }else{
                record[t]++;
            }
        }
        for (auto r : record){
            q.push(r.second);
        }
        int steps = 0;
        while (!q.empty()){
            int i = 0;
            std::priority_queue<int> q_bak;
            while (i < n + 1 && !q.empty()){
                int e = q.top();
                q.pop();
                e--;
                steps++;
                if(e > 0)
                    q_bak.push(e);
                i++;
                if (q.empty()){
                    if(!q_bak.empty())
                        steps += n  + 1 - i;
                    break;
                }
            }
            while (!q_bak.empty()){
                q.push(q_bak.top());
                q_bak.pop();
            }
        }
        return steps;
    }
};