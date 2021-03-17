#include <vector>
#include <unordered_map>
#include <queue>
using namespace std;

class Solution_347 {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        std::unordered_map<int,int> record;
        for(auto n: nums){
            record[n]++;
        }
        std::priority_queue<std::pair<int,int>> q;
        for(auto r : record){
            std::pair<int,int> p;
            p.first = r.second;
            p.second = r.first;
            q.push(p);
        }
        std::vector<int> res(k);
        int index = 0;
        while(k > 0){
            auto p = q.top();
            q.pop();
            res[index++] = p.second;
            k--;
        }
        return res;
    }
};