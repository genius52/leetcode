#include <vector>
#include <queue>
using namespace std;

class Solution_1383 {
public:
    int maxPerformance(int n, vector<int>& speed, vector<int>& efficiency, int k) {
        std::vector<std::pair<int,int>> record(n);
        for(int i = 0;i < n;i++)
            record[i] = {efficiency[i],speed[i]};
        std::sort(record.begin(),record.end(),[](std::pair<int,int>& p1,std::pair<int,int>& p2)->bool {
            if(p1.first == p2.first)
                return p1.second > p2.second;
            else
                return p1.first > p2.first;
        });
        std::priority_queue<int,std::vector<int>, std::greater<int>> q;//小顶堆,保存speed
        long res = 0;
        long sum_speed = 0;
        int MOD = 1e9 + 7;
        for(int i = 0;i < n;i++){
            sum_speed += record[i].second;
            if(i >= k){
                sum_speed -= q.top();
                q.pop();
            }
            q.push(record[i].second);
            res = max(res,sum_speed * record[i].first);
        }
        return res % MOD;
    }
};