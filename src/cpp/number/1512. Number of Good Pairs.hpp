#include <vector>
#include <unordered_map>
using namespace std;

class Solution_1512 {
public:
    int numIdenticalPairs(vector<int>& nums) {
        std::unordered_map<int,int> record;
        int len = nums.size();
        int res = 0;
        for(int i = 0;i < len;i++){
            record[nums[i]]++;
        }
        for(auto it = record.begin();it != record.end();it++){
            int cnt = (*it).second;
            if(cnt > 1){
                int total = 0;
                while(cnt-- > 1){
                    total += cnt;
                }
                res += total;
            }
        }
        return res;
    }
};