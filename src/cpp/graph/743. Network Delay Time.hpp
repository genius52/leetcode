#include <vector>
#include <unordered_map>
#include <queue>
using namespace std;

class Solution_743 {
public:
    int networkDelayTime(vector<vector<int>>& times, int n, int k) {
        std::vector<int> dis(n + 1,2147483647);
        dis[k] = 0;
        for(int i = 0;i < n;i++){
            for(auto t : times){
                auto cur_start = t[0];
                auto cur_end = t[1];
                auto duration = t[2];
                if(dis[cur_start] != 2147483647 && dis[cur_end] > dis[cur_start] + duration){
                    dis[cur_end] = dis[cur_start] + duration;
                }
            }
        }
        int max = 0;
        for(int i = 1;i <= n;i++){
            if(dis[i] == 2147483647)
                return -1;
            if (dis[i] > max)
                max = dis[i];
        }
        return max;
    }
};