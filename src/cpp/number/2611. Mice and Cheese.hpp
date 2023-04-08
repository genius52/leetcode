#include <vector>
#include <queue>
using namespace std;

class Solution_2611 {
public:
    int miceAndCheese(vector<int>& reward1, vector<int>& reward2, int k) {
        if(k == 0){
            return std::accumulate(reward2.begin(),reward2.end(),0);
        }
        int len = reward1.size();
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;
        std::vector<int> second(len - k);
        int j = 0;
        for(int i = 0;i < len;i++){
            if(q.size() < k){
                q.push({reward1[i] - reward2[i],i});
            }else{
                auto top = q.top();
                if((reward1[i] - reward2[i]) > top.first){
                    second[j++] = top.second;
                    q.pop();
                    q.push({reward1[i] - reward2[i],i});
                }else{
                    second[j++] = i;
                }
            }
        }
        int res = 0;
        while(!q.empty()){
            auto top = q.top();
            res += reward1[top.second];
            q.pop();
        }
        for(int i = 0;i < len - k;i++){
            res += reward2[second[i]];
        }
        return res;
    }
};