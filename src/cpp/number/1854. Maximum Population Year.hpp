#include <vector>
#include <set>
using namespace std;

class Solution_1854 {
public:
    int maximumPopulation(vector<vector<int>>& logs) {
        int len = logs.size();
        std::multiset<int> birth;
        std::multiset<int> death;
        int max_cnt = 0;
        int year = 0;
        for(int i = 0;i < len;i++){
            birth.insert(logs[i][0]);
            death.insert(logs[i][1]);
        }
        for(auto it = birth.begin();it != birth.end();it++){
            auto birthnum_before_current = std::distance(birth.begin(),it) + 1;//birth before current year
            auto deadnum_before_current = std::distance(death.begin(),death.upper_bound(*it));//death number before current year
            if(birthnum_before_current - deadnum_before_current > max_cnt){
                max_cnt = birthnum_before_current - deadnum_before_current ;
                year = *it;
            }
        }
        return year;
    }
};