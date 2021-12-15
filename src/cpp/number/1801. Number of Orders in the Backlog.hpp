#include <vector>
#include <queue>
using namespace std;


// orders[i] = [pricei, amounti, orderTypei]
//0 if it is a batch of buy orders, or
//1 if it is a batch of sell orders.
class Solution_1801 {
public:
    int getNumberOfBacklogOrders(vector<vector<int>>& orders) {
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::less<std::pair<int,int>>> buy;
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> sell;
        int len = orders.size();
        int MOD = 1e9 + 7;
        for (int i = 0; i < len; ++i) {
            if (orders[i][2] == 0){//buy,find lowest price in sell queue
                if (sell.empty()){
                    buy.push({orders[i][0],orders[i][1]});
                }else{
                    int cnt = orders[i][1];
                    while (!sell.empty() && cnt != 0){
                        auto top = sell.top();
                        if(sell.top().first > orders[i][0]){
                            break;
                        }
                        sell.pop();
                        if(cnt >= top.second){
                            cnt -= top.second;
                        }else{
                            top.second -= cnt;
                            cnt = 0;
                            sell.push(top);
                        }
                    }
                    if(cnt > 0){
                        buy.push({orders[i][0],cnt});
                    }
                }
            }else{//sell,find highest price in buy queue
                if (buy.empty()){
                    sell.push({orders[i][0],orders[i][1]});
                }else{
                    int cnt = orders[i][1];
                    while (!buy.empty() && cnt != 0){
                        auto top = buy.top();
                        if(buy.top().first < orders[i][0]){
                            break;
                        }
                        buy.pop();
                        if(cnt >= top.second){
                            cnt -= top.second;
                        }else{
                            top.second -= cnt;
                            cnt = 0;
                            buy.push(top);
                        }
                    }
                    if(cnt > 0){
                        sell.push({orders[i][0],cnt});
                    }
                }
            }
        }
        int res = 0;
        while(!buy.empty()){
            res += buy.top().second;
            res %= MOD;
            buy.pop();
        }
        while(!sell.empty()){
            res += sell.top().second;
            res %= MOD;
            sell.pop();
        }
        return res;
    }
};