#include <vector>
#include <queue>
#include <math.h>
using namespace std;

//TLE
class Solution_1792 {
public:
    double maxAverageRatio(vector<vector<int>>& classes, int extraStudents) {
        int len = classes.size();
        if(len == 1){
            return double(classes[0][0] + extraStudents)/double(classes[0][1] + extraStudents);
        }
        struct mycmp{
            bool operator()(const std::pair<double,std::vector<int>>& p1,const std::pair<double,std::vector<int>>& p2) {
                if (p1.first > p2.first)
                    return false;
                return true;
            }
        };
        std::priority_queue<std::pair<double,std::vector<int>>,std::vector<std::pair<double,std::vector<int>>>,mycmp> q;
        for(int i = 0;i < len;i++){
            double d1 = double(classes[i][0] + 1)/double(classes[i][1] + 1);
            double d2 = double(classes[i][0])/double(classes[i][1]);
            double profit = d1 - d2;
            std::pair<double,std::vector<int>> pair;
            pair.first = profit;
            pair.second = classes[i];
            //res += d2;
            q.push(pair);
        }
        while(extraStudents > 0){
            auto top = q.top();
            q.pop();
            double d1 = top.second[0] + 1;
            double d2 = top.second[1] + 1;
            auto profit = (d1 + 1)/(d2 + 1) - d1/d2;
            std::pair<double,std::vector<int>> pair;
            pair.first = profit;
            pair.second = std::vector<int>{top.second[0] + 1,top.second[1] + 1};
            q.push(pair);
            extraStudents--;
        }
        double res = 0.0;
        while(!q.empty()){
            auto top = q.top();
            q.pop();
            double d1 = top.second[0];
            double d2 = top.second[1];
            res += d1/d2;
        }
        return res/len;
    }
};