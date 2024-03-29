#include <vector>
#include <queue>
#include <map>
#include <unordered_set>
using namespace std;

class Solution_3092 {
public:
    vector<long long> mostFrequentIDs(vector<int>& nums, vector<int>& freq) {
        std::map<long long,std::unordered_set<int>> cnt_ids;
        int len = nums.size();
        std::unordered_map<int,long long> cur_cnt;
        for(int i = 0;i < len;i++)
            cur_cnt[nums[i]] = 0;
        std::vector<long long> res(len);
        for(int i = 0;i < len;i++){
            auto id = nums[i];
            auto change_cnt = freq[i];
            if(cur_cnt[id] != 0){
                cnt_ids[cur_cnt[id]].erase(id);
                if(cnt_ids[cur_cnt[id]].size() == 0){
                    cnt_ids.erase(cur_cnt[id]);
                }
            }
            cur_cnt[id] += change_cnt;
            cnt_ids[cur_cnt[id]].insert(id);
            if(!cnt_ids.empty())
                res[i] = cnt_ids.rbegin()->first;
        }
        return res;
    }
};