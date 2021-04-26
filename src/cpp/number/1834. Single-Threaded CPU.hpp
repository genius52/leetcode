#include <vector>
#include <map>
#include <set>
#include <unordered_set>
using namespace std;

//Input: tasks = [[1,2],[2,4],[3,2],[4,1]]
//Output: [0,2,3,1]
class Solution_1834 {
public:
    vector<int> getOrder(vector<vector<int>>& tasks) {
        //duration:index
        std::priority_queue<std::pair<int,int>,std::deque<std::pair<int,int>>,std::greater<std::pair<int,int>>>  q;
        int len = tasks.size();
        long start_time = ((long long)1<<31)-1;
        std::map<int,std::map<int,std::unordered_set<int>>> starttime_duration;//starttime:duration:index1,index2
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