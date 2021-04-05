#include <vector>
#include <queue>
using namespace std;

class Solution_378 {
public:
    int kthSmallest(vector<vector<int>>& matrix, int k) {
        std::priority_queue<int> q;
        for(auto m : matrix){
            for(auto n : m){
                if(q.size() < k){
                    q.push(n);
                }else{
                    int top = q.top();
                    if(n < top){
                        q.pop();
                        q.push(n);
                    }

                }
            }
        }
        return q.top();
    }
};