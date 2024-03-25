#include <vector>
#include <map>
#include <set>
using namespace std;

class Solution_3080 {
public:
    vector<long long> unmarkedSumArray(vector<int>& nums, vector<vector<int>>& queries) {
        int l1 = nums.size();
        int l2 = queries.size();
        std::vector<long long> res(l2);
        std::map<int,std::set<int>> data_index;
        std::vector<bool> visited(l1);//index
        long long sum = 0;
        for(int i = 0;i < l1;i++){
            sum += nums[i];
            data_index[nums[i]].insert(i);
        }
        for(int i = 0;i < l2;i++){
            int idx = queries[i][0];
            int k = queries[i][1];
            if(!visited[idx]){//如果没访问过，标记并删除字典里的索引
                visited[idx] = true;
                sum -= nums[idx];
                data_index[nums[idx]].erase(idx);
                if(data_index[nums[idx]].size() == 0){
                    data_index.erase(nums[idx]);
                }
            }
            if(data_index.size() == 0)
                break;
            while(k > 0){
                auto it = data_index.begin();
                sum -= it->first;
                visited[*it->second.begin()] = true;
                it->second.erase(it->second.begin());
                if(it->second.size() == 0){
                    data_index.erase(it);
                }
                if(data_index.size() == 0)
                    break;
                k--;
            }
            res[i] = sum;
            if(sum == 0)
                break;
        }
        return res;
    }
};