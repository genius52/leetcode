#include <vector>
#include <map>
using namespace std;

//输入：apples = [1,2,3,5,2], days = [3,2,1,4,2]
//输出：7
//解释：你可以吃掉 7 个苹果：
//- 第一天，你吃掉第一天长出来的苹果。
//- 第二天，你吃掉一个第二天长出来的苹果。
//- 第三天，你吃掉一个第二天长出来的苹果。过了这一天，第三天长出来的苹果就已经腐烂了。
//- 第四天到第七天，你吃的都是第四天长出来的苹果。
class Solution_1705 {
public:
    //The intuition is we should eat the apple with the closest expiration date
    int eatenApples(vector<int>& apples, vector<int>& days){
        int len = apples.size();
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>> q;//day,apples
        int res = 0;
        for(int i = 0;i < len || !q.empty();i++){
            while (!q.empty() && q.top().first <= i){
                q.pop();
            }
            if(i < len){
                if (apples[i] > 0)
                    q.push({i + days[i],apples[i]});
            }
            if(!q.empty()){
                auto p = q.top();
                q.pop();
                res++;
                p.second--;
                if(p.second > 0){
                    q.push(p);
                }
            }
        }
        return res;
    }

//    int eatenApples(vector<int>& apples, vector<int>& days) {
//        std::map<int,int> record;//key: date, val:count
//        int len = apples.size();
//        int total = 0;
//        int res = 0;
//        for(int i = 0;i < len;i++){
//            if(apples[i] != 0){
//                record[i + days[i] - 1] += apples[i];
//                total += apples[i];
//            }
//            if(total > 0){
//                auto it = record.lower_bound(i);
//                if(it->first == i || it->second == 1){
//                    total -= it->second;
//                    res++;
//                    record.erase(it);
//                }else{
//                    res++;
//                    it->second--;
//                    if(it->second == 0){
//                        record.erase(it);
//                    }
//                    total--;
//                }
//            }
//        }
//        int date = len - 1;
//        for(auto it = record.begin();it != record.end();it++){
//            int m = min(it->first - date,it->second);
//            res += m;
//            if(it->second >= (it->first - date)){
//                date = it->first;
//            }else{
//                date += (it->first - date - it->second);
//            }
//        }
//        return res;
//    }
};