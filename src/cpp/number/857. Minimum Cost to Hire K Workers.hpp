#include <queue>
#include <vector>
using namespace std;

//K个员工的价格=MAX(最低价格/能力) * SUM(能力)
class Solution_857 {
public:
    double mincostToHireWorkers(vector<int>& quality, vector<int>& wage, int k) {
        int len = quality.size();
        std::vector<std::pair<double,int>> rate_quality;
        //std::priority_queue<std::pair<double,std::pair<int,int>>> q;//薪水/能力，价性比
        for(int i = 0;i < len;i++){
            rate_quality.push_back({double(wage[i])/double(quality[i]),quality[i]});
        }
        std::sort(rate_quality.begin(),rate_quality.end());

        //总的支付工资：total wage = ratio * total quality
        std::priority_queue<int> q;
        double quality_sum = 0;
        double res = 2147483647;
        for(int i = 0;i < len;i++){//从单价便宜的开始找，同时记录总数
            quality_sum += rate_quality[i].second;
            q.push(rate_quality[i].second);
            if(q.size() > k){
                quality_sum -= q.top();//quality最大的，占用的薪水一定最高，所以干掉试试
                q.pop();
            }
            if(q.size() == k) {
                double cur_cost = rate_quality[i].first * quality_sum;
                res = min(res, cur_cost);
            }
        }
        return res;
    }
};