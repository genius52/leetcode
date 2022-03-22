#include <vector>
#include <queue>
using namespace std;

class Solution_2208 {
public:
    int halveArray(vector<int>& nums) {
        std::priority_queue<double, std::deque<double>, std::less<double> > q;
        double sum = 0;
        for(auto n : nums){
            sum += n;
            q.push(n);
        }
        double target = sum / 2;
        int steps = 0;
        while(sum > target){
            double val = q.top();
            sum -= val;
            q.pop();
            q.push(val / 2);
            sum += val / 2;
            steps++;
        }
        return steps;
    }
};