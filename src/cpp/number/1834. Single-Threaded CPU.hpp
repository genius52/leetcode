#include <vector>
#include <map>
#include <set>
#include <unordered_set>
#include <unordered_map>
#include <algorithm>
using namespace std;

//Input: tasks = [[1,2],[2,4],[3,2],[4,1]]
//Output: [0,2,3,1]
class Solution_1834 {
public:
    vector<int> getOrder2(vector<vector<int>>& tasks){
        int len = tasks.size();
        vector<vector<int>> record(len);
        for(int i = 0;i < len;i++){
            record[i] = {i,tasks[i][0],tasks[i][1]};//index,starttime,duration
        }
        std::sort(record.begin(),record.end(),[](const std::vector<int>& v1,const std::vector<int>& v2){
            if(v1[1] == v2[1]){
                return v1[2] < v2[2];
            }else{
                return v1[1] < v2[1];
            }
        });
        //返回true时，说明p1的优先级低于p2
        //duration值较大的Node优先级低（duration小的Node排在队前）
        //duration相等时，index大的优先级低（index小的Node排在队前）
        auto my_cmp = [](const std::pair<int,int>& p1,const std::pair<int,int>& p2)->bool {
            if(p1.first == p2.first){
                return p1.second > p2.second;
            }else{
                return p1.first > p2.first;
            }
        };
        //duration,index
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>, decltype(my_cmp)> wait_q(my_cmp);
        std::vector<int> res;
        long last_finish = 0;//上一个任务的结束时间
        int i = 0;
        while (i < len || !wait_q.empty()) {
            if(wait_q.empty()){
                last_finish = max(int(last_finish),record[i][1]);
            }
            while(i < len && record[i][1] <= last_finish){
                wait_q.push({record[i][2],record[i][0]});
                i++;
            }
            auto cur = wait_q.top();
            res.push_back(cur.second);
            last_finish += record[cur.second][2];
            wait_q.pop();
        }
        return res;
    }

    vector<int> getOrder(vector<vector<int>>& tasks) {
        //duration:index
        std::priority_queue<std::pair<int,int>,std::deque<std::pair<int,int>>,std::greater<std::pair<int,int>>>  q;
        int len = tasks.size();
        long start_time = ((long long)1<<31)-1;
        std::map<int,std::unordered_map<int,std::unordered_set<int>>> starttime_duration;//starttime:duration:index1,index2
        for(int i = 0;i < len;i++){
            if(tasks[i][0] < start_time){
                start_time = tasks[i][0];
            }
            starttime_duration[tasks[i][0]][tasks[i][1]].insert(i);
        }
        std::vector<bool> visited(len);
        std::vector<int> res(len);
        int index = 0;
        while(index < len){
            for(auto it = starttime_duration.begin();it != starttime_duration.end();){
                if(it->first <= start_time){
                    for(auto it2 = it->second.begin();it2 != it->second.end();it2++){
                        for(auto it3 = it2->second.begin();it3 != it2->second.end();it3++){
                            std::pair<int,int> p1;
                            p1.first = it2->first;
                            p1.second = *it3;
                            q.push(p1);
                        }
                    }
                    starttime_duration.erase(it++);
                }else{
                    break;
                }
            }
            if(q.empty()){
                start_time = starttime_duration.begin()->first;
                for(auto it = starttime_duration[start_time].begin();it != starttime_duration[start_time].end();it++){
                    for(auto it2 = it->second.begin();it2 != it->second.end();it2++){
                        std::pair<int,int> p1;
                        p1.first = it->first;
                        p1.second = *it2;
                        q.push(p1);
                    }
                }
                starttime_duration.erase(starttime_duration.begin());
            }else{
                auto p = q.top();
                q.pop();
                res[index] = p.second;
                start_time += p.first;
                index++;
            }
        }
        return res;
    }
};