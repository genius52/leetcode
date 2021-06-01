#include <vector>
#include <list>
#include <map>
#include <queue>
using namespace std;

class Solution_1882 {
public:
    vector<int> assignTasks(vector<int>& servers, vector<int>& tasks) {
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> free_q;//weight : index
        int ser_len = servers.size();
        for(int i = 0;i < ser_len;i++){
            free_q.push({servers[i],i});
        }
        std::map<int,std::list<std::pair<int,int>>> busy_q;//timstamp: weight,index
        int seconds = tasks.size();
        std::vector<int> res(seconds);
        int cur_time = 0;
        //        std::vector<int> servers{31,96,73,90,15,11,1,90,72,9,30,88};
        //        std::vector<int> tasks{87,10,3,5,76,74,38,64,16,64,93,95,60,79,54,26,30,44,64,71};
        for(int i = 0;i < seconds;){
            auto it_timestamp = busy_q.upper_bound(cur_time);//move busy machine to free queue before current timestamp
            if(it_timestamp != busy_q.begin()){
                for(auto it = busy_q.begin();it != it_timestamp;){
                    for(auto it2 = (*it).second.begin();it2 != (*it).second.end();it2++){
                        free_q.push(*it2);
                    }
                    busy_q.erase(it++);
                }
            }

            if(!free_q.empty()){
                std::pair<int,int> cur = free_q.top();
                free_q.pop();
                res[i] = cur.second;
                busy_q[cur_time + tasks[i]].push_back(cur);
                cur_time++;
                i++;
            }else{
                //get a machine from begining of busy queue then set into next timestamp
//                auto it = busy_q.begin();
//
//                cur_time = max(cur_time,(*it).first);
//                auto find = (*it).second.begin();
//                std::pair<int,int> p;
//                p.first = (*find).first;
//                p.second = (*find).second;
//                res[i] = (*find).second;
//                (*it).second.erase(find);
//                if((*it).second.size() == 0){
//                    busy_q.erase(it);
//                }
//                busy_q[cur_time + tasks[i]].push_back(p);
                auto it = busy_q.begin();
                cur_time = max(cur_time,(*it).first);
            }
        }
        return res;
    }
};
