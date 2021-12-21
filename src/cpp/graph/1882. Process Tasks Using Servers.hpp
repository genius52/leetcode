#include <vector>
#include <list>
#include <map>
#include <queue>
using namespace std;

class Solution_1882 {
public:
    vector<int> assignTasks(vector<int>& servers, vector<int>& tasks) {
        auto my_cmp = [](const std::pair<int,int>& p1,const std::pair<int,int>& p2)->bool {
            if(p1.first == p2.first){
                return p1.second > p2.second;
            }else{
                return p1.first > p2.first;
            }
        };
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>, decltype(my_cmp)>  free_server(my_cmp);//weight : index
        auto my_cmp2 = [](const std::pair<int,std::pair<int,int>>& p1,const std::pair<int,std::pair<int,int>>& p2)->bool {
            if(p1.first == p2.first){
                if(p1.second.first == p2.second.first)
                    return p1.second.second > p2.second.second;
                else
                    return p1.second.first > p2.second.first;
            }else{
                return p1.first > p2.first;
            }
        };
        std::priority_queue<std::pair<int,std::pair<int,int>>,std::vector<std::pair<int,std::pair<int,int>>>, decltype(my_cmp2)> busy_server(my_cmp2);//timestamp:server_weight,server_index;
        std::list<std::pair<int,int>> wait_task;//index,duration
        int ser_len = servers.size();
        for(int i = 0;i < ser_len;i++){
            free_server.push({servers[i],i});
        }
        int task_len = tasks.size();
        std::vector<int> res(task_len);
        for(int i = 0;i < task_len;i++){
            while(!busy_server.empty() && busy_server.top().first <= i){
                free_server.push(busy_server.top().second);
                busy_server.pop();
            }
            if(free_server.empty()){
                auto top = busy_server.top();
                busy_server.pop();
                res[i] = top.second.second;
                busy_server.push({top.first + tasks[i],top.second});
            }else{
                auto top = free_server.top();
                free_server.pop();
                res[i] = top.second;
                busy_server.push({i + tasks[i],top});
            }
        }
        return res;
    }
};
