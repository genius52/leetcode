#include <vector>
#include <queue>
using namespace std;

class Solution_1851 {
public:
    vector<int> minInterval(vector<vector<int>>& intervals, vector<int>& queries) {
        int len = queries.size();
        std::vector<int> res(len);
        std::vector<std::vector<int>> record(len);
        for (int i = 0; i < len; ++i) {
            record[i] = {queries[i],i};
        }
        std::sort(record.begin(),record.end());
        std::sort(intervals.begin(),intervals.end(),[](const std::vector<int>& v1,const std::vector<int>& v2){
            return v1[1] < v2[1];
        });
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        int interval_len = intervals.size();
        int idx = interval_len - 1;
        for (int i = len - 1; i >= 0; i--) {
            while(idx >= 0 && intervals[idx][1] >= record[i][0]){
                q.push({intervals[idx][1] - intervals[idx][0] + 1,idx});
                idx--;
            }
            if(q.empty()){
                res[record[i][1]] = -1;
            }else{
                while (!q.empty() && (intervals[q.top().second][0] > record[i][0] || intervals[q.top().second][1] < record[i][0])){
                    q.pop();
                }
                if(q.empty()){
                    res[record[i][1]] = -1;
                }else{
                    res[record[i][1]] = q.top().first;
                }
            }
        }
        return res;
    }
};