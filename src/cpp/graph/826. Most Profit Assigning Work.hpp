#include <vector>
#include <queue>
#include <set>
using namespace std;

class Solution_826 {
public:
    int maxProfitAssignment(vector<int>& difficulty, vector<int>& profit, vector<int>& worker) {
        int len = difficulty.size();
        std::priority_queue<std::pair<int,int>> q_job;
        for(int i = 0;i < len;i++){
            q_job.push({profit[i],difficulty[i]});
        }
        std::multiset<int> worker_set;
        for(auto w : worker){
            worker_set.insert(w);
        }
        int res = 0;
        while(worker_set.size() > 0 && q_job.size() > 0){
            auto cur = q_job.top();
            q_job.pop();
            auto it = worker_set.lower_bound(cur.second);//找到能力值能满足当前任务的工人,一个工作可以安排给多个工人
            if(it != worker_set.end()){
                auto cnt = std::distance(it,worker_set.end());
                res += cur.first * cnt;
                while(it != worker_set.end()){
                    worker_set.erase(it++);
                }
            }
        }
        return res;
    }
};