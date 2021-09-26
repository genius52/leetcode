#include <vector>
#include <queue>
//#include <algorithm>
#include <numeric>
using namespace std;

class Solution_1354 {
public:
    bool isPossible(vector<int>& target) {
        int len = target.size();
        if(len == 1)
            return target[0] == 1;
        long long sum = std::accumulate(target.begin(),target.end(),(long long)(0));
        std::priority_queue<long long> q(target.begin(),target.end());
        while (true){
            long long biggest = q.top();
            sum -= biggest;
            q.pop();
            if(biggest == 1 || sum == 1)//[1,i] always true
                return true;
            if(biggest < sum)
                return false;
            long long pre = biggest % sum;
            if(pre == 0)
                return false;
            q.push(pre);
            sum += pre;
        }
        return true;
    }
};