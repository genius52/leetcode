#include <vector>
#include <map>
#include <set>
#include <queue>
using namespace std;

//输入：k = 3, arrival = [1,2,3,4,5], load = [5,2,3,3,3]
//输出：[1]
//解释：
//所有服务器一开始都是空闲的。
//前 3 个请求分别由前 3 台服务器依次处理。
//请求 3 进来的时候，服务器 0 被占据，所以它呗安排到下一台空闲的服务器，也就是服务器 1 。
//请求 4 进来的时候，由于所有服务器都被占据，该请求被舍弃。
//服务器 0 和 2 分别都处理了一个请求，服务器 1 处理了两个请求。所以服务器 1 是最忙的服务器。
class Solution_1606 {
public:
    vector<int> busiestServers(int k, vector<int>& arrival, vector<int>& load) {
        std::map<int,std::set<int>> busy;//结束时间:号码集合
        std::set<int> free;
        std::vector<int> count(k);
        std::priority_queue<std::pair<int,int>> q;
        for(int i = 0;i < k;i++){
            free.insert(i);
        }
        int cur_idx = 0;
        for (int i = 0; i < arrival.size(); ++i) {
            cur_idx = i % k;
            int task_time = arrival[i];
            auto end = busy.upper_bound(task_time);
            for (auto it = busy.begin(); it != end;) { //把到期的任务都放入空闲队列
                for(auto it2 = it->second.begin();it2 != it->second.end();it2++){
                    free.insert(*it2);
                }
                busy.erase(it++);
            }
            if (free.empty()){//任务来了，没有空闲机器,被舍弃
                continue;
            }else{
                auto find = free.lower_bound(cur_idx);//寻找合适的服务器
                if(find == free.end()){
                    find = free.lower_bound(0);
                }
                cur_idx = *find;
                free.erase(find);
                busy[task_time + load[i]].insert(cur_idx);
                count[cur_idx]++;
            }
            //cur_idx = (cur_idx + 1) % k;
        }
        std::vector<int> res;
        int max_cnt = 0;
        for (int i = 0; i < k; ++i) {
            if(count[i] > max_cnt)
                max_cnt = count[i];
        }
        for (int i = 0; i < k; ++i){
            if(count[i] == max_cnt)
                res.push_back(i);
        }
        return res;
    }
};