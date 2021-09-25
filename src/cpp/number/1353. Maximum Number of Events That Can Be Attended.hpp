#include <vector>
#include <queue>
using namespace std;

class Solution_1353 {
public:
    int maxEvents(vector<vector<int>>& events) {
        int len = events.size();
        if(len <= 1)
            return len;
        int begin = 1000000;
        int end = 0;
        std::sort(events.begin(),events.end(),[&](const std::vector<int>& v1,const std::vector<int>& v2)->bool{
            if(v1[0] < begin)
                begin = v1[0];
            if(v2[0] < begin)
                begin = v2[0];
            if(v1[1] > end)
                end = v1[1];
            if(v2[1] > end)
                end = v2[1];
            if(v1[0] == v2[0]){
                return v1[1] < v2[1];
            }
            return v1[0] < v2[0];
        });
        std::priority_queue<std::pair<int,int>,std::deque<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        int cnt = 0;
        int idx = 0;
        //1.将以当前结束时间开始的event加入队列，
        //2.选中队列中最早结束的事件，
        //3.当前时间+1
        //4.删除队列中 结束时间小于当前时间的事件
        while (idx < len || begin <= end){
            while (idx < len && events[idx][0] == begin){
                q.push({events[idx][1],events[idx][0]});
                idx++;
            }
            if(!q.empty()){
                std::pair<int,int> cur = q.top();
                q.pop();
                cnt++;
            }
            begin++;
            while (!q.empty() && q.top().first < begin){
                q.pop();
            }
        }
        return cnt;
    }
};