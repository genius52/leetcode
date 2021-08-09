#include <vector>
#include <map>
using namespace std;

class Solution_1964 {
public:
    vector<int> longestObstacleCourseAtEachPosition(vector<int>& obstacles) {
        int len = obstacles.size();
        std::vector<int> len_num;
        len_num.push_back(obstacles[0]);
        std::vector<int> res(len);
        res[0] = 1;
        for(int i = 1;i < len;i++){
            auto it = std::upper_bound(len_num.begin(),len_num.end(),obstacles[i]);
            if(it == len_num.end()){
                len_num.push_back(obstacles[i]);
                res[i] = len_num.size();
            }else{
                *it = obstacles[i];
                res[i] = std::distance(len_num.begin(),it) + 1;
            }
        }
        return res;
    }
};