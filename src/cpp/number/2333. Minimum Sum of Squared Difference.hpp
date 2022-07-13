#include <vector>
#include <queue>

class Solution_2333 {
public:
    long long minSumSquareDiff(vector<int>& nums1, vector<int>& nums2, int k1, int k2) {
        int len = nums1.size();
        std::unordered_map<int,int> diff_cnt;
        long long total_diff = 0;
        int max_diff = 0;
        for(int i = 0;i < len;i++){
            int diff = abs(nums1[i] - nums2[i]);
            if(diff == 0){
                continue;
            }
            if(diff > max_diff)
                max_diff = diff;
            total_diff += diff;
            diff_cnt[diff]++;
        }
        int k = k1 + k2;
        if(k >= total_diff)
            return 0;
        std::vector<int> record(max_diff + 1);
        for(auto it : diff_cnt)
            record[it.first] = it.second;
        for(int i = max_diff;i > 0;i--){
            int decrease = min(k,record[i]);
            record[i] -= decrease;
            record[i - 1] += decrease;
            k -= decrease;
            if(k == 0)
                break;
        }
        long long res = 0;
        for(long long i = max_diff;i > 0;i--){
            if(record[i] == 0)
                continue;
            res += i *  i * record[i];
        }
        return res;
    }

//    long long minSumSquareDiff(vector<int>& nums1, vector<int>& nums2, int k1, int k2) {
//        int len = nums1.size();
//        std::unordered_map<int,int> diff_cnt;
//        int total_diff = 0;
//        for(int i = 0;i < len;i++){
//            int diff = nums1[i] - nums2[i];
//            if(diff == 0){
//                continue;
//            }
//            if(diff < 0){
//                diff = -diff;
//            }
//            total_diff += diff;
//            diff_cnt[diff]++;
//        }
//        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::less<std::pair<int,int>>> q;//big top
//        for(auto it : diff_cnt){
//            q.push({it.first,it.second});
//        }
//        int k = k1 + k2;
//        if(k >= total_diff)
//            return 0;
//        while(k != 0 && !q.empty()){
//            int diff1 = q.top().first;
//            int cnt1 = q.top().second;
//            q.pop();
//            if(!q.empty()){
//                int diff2 = q.top().first;
//                int target = (diff1 - diff2 + 1) * cnt1;
//                if(k >= target){
//                    diff1 -= target;
//                    if(diff1 != 0)
//                        q.push({diff1,cnt1});
//                    k -= target;
//                }else{
//                    if(k < cnt1){
//                        q.push({diff1 - 1,k});
//                        q.push({diff1,cnt1 - k});
//                        k = 0;
//                    }else{
//                        if(diff1 - k/cnt1 > 0){
//                            q.push({diff1 - k/cnt1,cnt1 - k % cnt1});
//                        }
//                        if(diff1 - k /cnt1 - 1 > 0){
//                            q.push({diff1 - k /cnt1 - 1,k % cnt1});
//                        }
//                    }
//                    k = 0;
//                }
//                if(k == 0)
//                    break;
//            }else{
//                int target = diff1 * cnt1;
//                if(k < target){
//                    if(k < cnt1){
//                        if(diff1 - 1 > 0){
//                            q.push({diff1 - 1,k});
//                        }
//                        q.push({diff1,cnt1 - k});
//                    }else{
//                        if(diff1 - k/cnt1 > 0){
//                            q.push({diff1 - k/cnt1,cnt1 - k % cnt1});
//                        }
//                        if(diff1 - k /cnt1 - 1 > 0){
//                            q.push({diff1 - k /cnt1 - 1,k % cnt1});
//                        }
//                    }
//                    k = 0;
//                }
//                break;
//            }
//        }
//        long long res = 0;
//        while(!q.empty()){
//            res += (long long)q.top().first * (long long)q.top().first * (long long)q.top().second;
//            q.pop();
//        }
//        return res;
//    }
};