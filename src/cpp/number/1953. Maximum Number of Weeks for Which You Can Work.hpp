#include <vector>
using namespace std;

class Solution_1953 {
public:
    long long numberOfWeeks(vector<int>& milestones){
        int len = milestones.size();
        long long sum = 0;
        long long max_val = 0;
        for(int i = 0;i < len;i++){
            sum += milestones[i];
            if(milestones[i] > max_val)
                max_val = milestones[i];
        }
        auto rest = sum - max_val;
        if(max_val > (rest + 1)){
            return rest * 2 + 1;
        }
        return sum;
    }
//    long long numberOfWeeks(vector<int>& milestones) {
//        std::priority_queue<int> q;
//        int len = milestones.size();
//        int sum = 0;
//        for(int i = 0;i < len;i++){
//            sum += milestones[i];
//            q.push(milestones[i]);
//        }
//        int res = 0;
//        while(q.size() > 1){
//            auto first = q.top();
//        }
//        if(q.size() > 0){
//            res++;
//        }
//        return res;
//    }
};